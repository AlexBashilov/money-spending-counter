package postgres

import (
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog"
)

const (
	// Значения по умолчанию для конфигурации подключения.
	defaultMaxOpenConnections = 0
	defaultDialTimeout        = 5 * time.Second
	defaultSlowQueryThreshold = 500 * time.Millisecond
)

// config - это структура конфигурации для установки корректного
// подключения к базе данных через данную библиотеку.
// Это отдельная сущность, которая реализует наиболее важные и используемые
// в сервисах параметры.
type config struct {
	// dsn - строка подключения к базе данных.
	// Обязательный параметр. Не имеет значения по умолчанию.
	dsn string

	// driverType - драйвер, используемый для создания Connector.
	// Выбор драйвера реализован константами PGXDriver и BunPGDriver,
	// Необязательный параметр. Значение по умолчанию: PGXDriver.
	driverType DriverType

	// networkType - тип подключения к базе данных.
	// Необязательный параметр. Имеет значение по умолчанию (tcp4).
	networkType string

	// appName - имя сервиса, используемого для идентификации подключения.
	// Обязательный параметр. Не имеет значения по умолчанию.
	appName string

	// maxOpenConnections - максимальное количество одновременно открытых соединений
	// с базой данных. Если значение равно 0, то нет ограничений.
	// Имеет значение по умолчанию (0).
	maxOpenConnections int

	// maxIdleConnections - максимальное количество соединений в режиме ожидания.
	// Если значение равно 0, то нет ограничений.
	// Необязательный параметр.
	maxIdleConnections int

	// dialTimeout - таймаут на установку соединения с базой данных.
	// Имеет значение по умолчанию (5 секунд).
	dialTimeout time.Duration

	// databaseName - название базы данных.
	// Подставляется автоматически из DSN.
	databaseName string

	// validateConnFunc - функция валидации соединения.
	// Имеет значение по умолчанию (pgconn.ValidateConnectTargetSessionAttrsReadWrite).
	// Необязательный параметр.
	validateConnFunc pgconn.ValidateConnectFunc

	// queryExecMode - режим общения с БД.
	// Значение по-умолчанию определяется драйвером. На данный момент pgx.QueryExecModeCacheStatement.
	// Необязательный параметр.
	queryExecMode *pgx.QueryExecMode

	// slowQueryThreshold - порог длительности для определения медленных запросов.
	// Значение по умолчанию `500 * time.Millisecond`.
	// Необязательный параметр.
	slowQueryThreshold time.Duration

	// slowQueryLogLevel - уровень логирования для медленных запросов.
	// Значение по умолчанию `zerolog.Disabled`.
	// Необязательный параметр.
	slowQueryLogLevel zerolog.Level

	// queryLogLevel - уровень логирования для общих запросов.
	// Значение по умолчанию `zerolog.Disabled`.
	// Необязательный параметр.
	queryLogLevel zerolog.Level

	// queryErrorLogLevel - уровень логирования для запросов, которые завершились с ошибкой.
	// Значение по умолчанию `zerolog.ErrorLevel`.
	// Необязательный параметр.
	queryErrorLogLevel zerolog.Level

	// readOnly - флаг, определяющий тип подключения к базе данных: реплика или мастер.
	// Значение по умолчанию `false`.
	// Необязательный параметр.
	readOnly bool

	// withAdditionalSuffix - флаг, который позволяет проставлять дополнительный суффикс к
	// имени базы данных, позволяющий различать при мониторинге мастер/реплику.
	// Значение по умолчанию `false`.
	// Необязательный параметр.
	withAdditionalSuffix bool

	// traceEnabled - флаг, который позволяет отключить трейсинг запросов
	// Значение по умолчанию `true`.
	// Необязательный параметр.
	traceEnabled bool
}

// NewConfig возвращает конфигурацию для создания коннектора и подключения к базе данных.
// Конфигурация имеет обязательные поля, а также возможность добавить некоторые
// параметры с помощью хелперов.
func newConfig(dsn, name string, opts ...Option) config {
	cfg := config{
		dsn:                  dsn,
		databaseName:         "", // будет заполнено автоматически
		driverType:           PGXDriver,
		networkType:          dialerNetwork,
		appName:              name,
		maxOpenConnections:   defaultMaxOpenConnections,
		maxIdleConnections:   0,
		dialTimeout:          defaultDialTimeout,
		validateConnFunc:     pgconn.ValidateConnectTargetSessionAttrsReadWrite,
		queryExecMode:        nil,
		slowQueryThreshold:   defaultSlowQueryThreshold,
		slowQueryLogLevel:    zerolog.Disabled,
		queryLogLevel:        zerolog.Disabled,
		queryErrorLogLevel:   zerolog.ErrorLevel,
		readOnly:             false,
		withAdditionalSuffix: false,
		traceEnabled:         true,
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	if cfg.readOnly {
		cfg.validateConnFunc = pgconn.ValidateConnectTargetSessionAttrsReadOnly
	}

	return cfg
}

// configValidation Валидирует конфигурацию, проверяя наличие значений в обязательных полях:
//
//	dsn - строка подключения.
//	appName - имя сервиса.
//
// Возвращает ошибку, если обязательные поля имеют значения по умолчанию.
func (c *config) configValidation() error {
	if c.dsn == "" {
		return ErrEmptyDSN
	}

	if c.appName == "" {
		return ErrEmptyAppName
	}

	return nil
}
