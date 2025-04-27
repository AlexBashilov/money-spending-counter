// postgres пакет создан для унификации подключения к базе данных postgresql внутри компании. Помимо
// инициализации имеется возможность настройки подключения с помощью обязательных и необязательных параметров,
// настраиваемых через хелперы.
//
// На данный момент используется резолв только IPv4 адресов.
// Из коробки реализуется сбор метрик и трейсов по запросам в базу с помощью OpenTelemetry.

package postgres

import (
	"context"
	"database/sql/driver"
	"net"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/newrelic/go-agent/v3/newrelic/sqlparse"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bunotel"
	"github.com/uptrace/bun/extra/bunzerolog"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	"go.opentelemetry.io/otel/trace/noop"

	"booker/pkg/postgres/reconnect"
)

// DriverType - драйвер, используемый при создании driver.Connector.
type DriverType int

const (
	// Использовать pgx для создания driver.Connector.
	PGXDriver DriverType = iota + 1
	// Использовать bun/pgdriver для создания driver.Connector.
	BunPGDriver
)

const (
	dialerKeepAlive = 5 * time.Minute
	// Используется для резолва ip-адресов только версии IPv4, исключая IPv6.
	dialerNetwork = "tcp4"

	primaryConnection = "-primary"
	replicaConnection = "-replica"

	// Ограничиваем время жизни коннектов к БД.
	connMaxLifetime = 30 * time.Minute
)

// NewConnection возвращает экземпляр структуры подключения к базе данных, обернутый в ORM Bun - *bun.DB.
// Точка входа для работы с библиотекой.
//
// Использует для создания подключения к postgres коннектор - driver.Connector.
// По умолчанию оборачивает экземпляр *bun.DB в обертку open telemetry для дальнейшей передачи дефолтных метрик
// и трейсов при настроенных экспортерах.
func NewConnection(dsn, name string, opts ...Option) (*bun.DB, error) {
	cfg := newConfig(dsn, name, opts...)

	if err := cfg.configValidation(); err != nil {
		return nil, errors.Wrap(err, "postgres: validating config")
	}

	connector, err := prepareDBConnector(&cfg)
	if err != nil {
		return nil, errors.Wrap(err, "prepare db connector")
	}

	return createConnection(&cfg, connector), nil
}

// prepareDBConnector - создает driver.Connector для подключения к базе данных,
// с учётом выбранного драйвера и прочих параметров конфигурации.
func prepareDBConnector(cfg *config) (driver.Connector, error) {
	if cfg.driverType == BunPGDriver {
		return prepareBunConnector(cfg)
	}

	return preparePGXConnector(cfg)
}

// prepareBunConnector создает и настраивает pgdriver.Connector для подключения к базе данных,
// основываясь на переданных параметрах конфигурации.
//
// Параметры:
//   - cfg: конфигурация подключения к базе данных.
//
// Результат:
//   - driver.Connector: реализация коннектора (pgdriver.Connector) для подключения к базе данных.
//   - error: ошибка при создании коннектора.
func prepareBunConnector(cfg *config) (driver.Connector, error) {
	// Создаем и настраиваем коннектор для подключения к базе данных.
	connector := pgdriver.NewConnector(
		pgdriver.WithDSN(cfg.dsn),
		pgdriver.WithApplicationName(cfg.appName),
		// Установить tcp4 без хардкода в Dialer при выполнении подключения.
		pgdriver.WithNetwork(dialerNetwork),
		// Определить Dial-, Read- и Write- таймауты.
		pgdriver.WithTimeout(cfg.dialTimeout),
		func(pgCfg *pgdriver.Config) {
			pgCfg.Dialer = func(ctx context.Context, network, addr string) (net.Conn, error) {
				return (&net.Dialer{ //nolint:exhaustruct
					FallbackDelay: -1,
					Timeout:       cfg.dialTimeout,
					KeepAlive:     dialerKeepAlive,
				}).DialContext(ctx, network, addr)
			}
		},
	)

	// Устанавливаем имя базы данных.
	cfg.databaseName = getDatabaseName(cfg, connector.Config().Database)

	return connector, nil
}

// preparePGXConnector создает коннектор для подключения к базе данных,
// с использованием библиотеки pgx, основываясь на переданных параметрах конфигурации.
//
// Parameters:
//   - cfg: конфигурация подключения к базе данных.
//
// Returns:
//   - driver.Connector: коннектор для подключения к базе данных.
//   - error: ошибка при создании коннектора.
func preparePGXConnector(cfg *config) (driver.Connector, error) {
	// Разбираем конфигурацию подключения к базе данных.
	connConfig, err := pgx.ParseConfig(cfg.dsn)
	if err != nil {
		return nil, errors.Wrap(err, "parse config")
	}

	// Обычно используется для отмены prepared statement
	// Больше информации https://pkg.go.dev/github.com/jackc/pgx/v5#QueryExecMode
	if cfg.queryExecMode != nil {
		connConfig.DefaultQueryExecMode = *cfg.queryExecMode
	}

	// Константа аттрибута target_session_attrs из libpq.
	// https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNECT-TARGET-SESSION-ATTRS
	const targetSessionAttrs = "target_session_attrs"

	// Если была предоставлена функция валидации подключения, устанавливаем ее в конфигурацию подключения.
	// Валидация не будет выполняться, если в DSN указан атрибут target_session_attrs.
	// По умолчанию используется pgconn.ValidateConnectTargetSessionAttrsReadWrite.
	if cfg.validateConnFunc != nil && !strings.Contains(cfg.dsn, targetSessionAttrs) {
		connConfig.ValidateConnect = cfg.validateConnFunc
	}

	// Заменяем стандартную функцию для создания соединения.
	connConfig.DialFunc = func(ctx context.Context, _, addr string) (net.Conn, error) {
		return (&net.Dialer{ //nolint:exhaustruct
			FallbackDelay: -1,
			Timeout:       cfg.dialTimeout,
			KeepAlive:     dialerKeepAlive,
		}).DialContext(ctx, cfg.networkType, addr)
	}
	connConfig.ConnectTimeout = cfg.dialTimeout
	connConfig.RuntimeParams["application_name"] = cfg.appName

	// Устанавливаем имя базы данных.
	cfg.databaseName = getDatabaseName(cfg, connConfig.Database)

	// Создаем коннектор для подключения к базе данных.
	connector := stdlib.GetConnector(*connConfig)

	return connector, nil
}

// getDatabaseName - конструктор, собирающий имя базы данных в зависимости от требований пользователя.
func getDatabaseName(cfg *config, database string) string {
	if !cfg.withAdditionalSuffix {
		return database
	}

	if cfg.readOnly {
		return database + replicaConnection
	}

	return database + primaryConnection
}

// createConnection непосредственно создает экземпляр подключения к базе данных через ORM Bun с использованием
// библиотеки open telemetry для дальнейшей работы с метриками и трейсами запросов.
//
//nolint:exhaustruct
func createConnection(cfg *config, connector driver.Connector) *bun.DB {
	otelsqlOptions := []otelsql.Option{}
	if !cfg.traceEnabled {
		otelsqlOptions = append(otelsqlOptions, otelsql.WithTracerProvider(noop.NewTracerProvider()))
	}

	odb := otelsql.OpenDB(newrelic.InstrumentSQLConnector(connector, newrelic.SQLDriverSegmentBuilder{
		BaseSegment: newrelic.DatastoreSegment{
			Product:      newrelic.DatastorePostgres,
			DatabaseName: cfg.databaseName,
		},
		ParseQuery: func(segment *newrelic.DatastoreSegment, query string) {
			sqlparse.ParseQuery(segment, query)

			segment.Operation = strings.ToUpper(segment.Operation)
			segment.ParameterizedQuery = query
			segment.RawQuery = query
		},
	}), otelsqlOptions...)

	if cfg.maxIdleConnections > 0 {
		odb.SetMaxIdleConns(cfg.maxIdleConnections)
	}

	odb.SetMaxOpenConns(cfg.maxOpenConnections)

	// bun.NewDB создает экземпляр Bun db.
	// В качестве опции используем WithDiscardUnknownColumns для игнорирования неизвестных столбцов,
	// больше информации https://bun.uptrace.dev/guide/running-bun-in-production.html#bun-withdiscardunknowncolumns
	conn := bun.NewDB(odb, pgdialect.New(), bun.WithDiscardUnknownColumns())

	bunotelOptions := []bunotel.Option{bunotel.WithDBName(cfg.databaseName)}
	if !cfg.traceEnabled {
		bunotelOptions = append(bunotelOptions, bunotel.WithTracerProvider(noop.NewTracerProvider()))
	}

	conn.AddQueryHook(bunotel.NewQueryHook(bunotelOptions...))

	conn.AddQueryHook(bunzerolog.NewQueryHook(
		bunzerolog.WithSlowQueryLogLevel(cfg.slowQueryLogLevel),
		bunzerolog.WithSlowQueryThreshold(cfg.slowQueryThreshold),
		bunzerolog.WithErrorQueryLogLevel(cfg.queryErrorLogLevel),
		bunzerolog.WithQueryLogLevel(cfg.queryLogLevel),
	))

	// Для RW-соединений, при использовании bun/pgdriver, добавляем проверку на смену мастера.
	if !cfg.readOnly && cfg.driverType == BunPGDriver {
		hook := reconnect.NewQueryHook()
		conn.AddQueryHook(hook.WithConnLifetimeNormal(connMaxLifetime))
	}

	return conn
}
