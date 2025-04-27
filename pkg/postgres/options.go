package postgres

import (
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
)

// Option - тип опциональных параметров конфигурации.
type Option func(cfg *config)

// WithOpenConnections возвращает замыкание для замены значения по умолчанию
// maxOpenConnections на требуемое разработчику.
//
// Параметры:
//
//	c - количество открытых соединений, которые можно установить.
//
// Пример:
//
//	postgres.NewConnection(
//		"postgres://root@localhost:5432/?sslmode=disable",
//		"app-name",
//		postgres.WithOpenConnections(10),
//	)
func WithOpenConnections(c int) Option {
	return func(cfg *config) {
		cfg.maxOpenConnections = c
	}
}

// WithDialTimeout возвращает замыкание для замены значения по умолчанию
// dialTimeout на требуемое разработчику.
//
// Параметры:
//
//	t - время ожидания установки соединения с сервером.
//
// Результат:
//
//	Option - функция, которая применяет переданное значение к конфигурации.
//
// Пример:
//
//	postgres.NewConnection(
//		"postgres://root@localhost:5432/?sslmode=disable",
//		"app-name",
//		postgres.WithDialTimeout(10 * time.Second),
//	)
func WithDialTimeout(t time.Duration) Option {
	return func(cfg *config) {
		cfg.dialTimeout = t
	}
}

// WithMaxIdleConnections возвращает замыкание, которое применяется к конфигурации.
// Это замыкание устанавливает значение максимального количества готовых для использования
// соединений с базой данных.
//
// Параметры:
//
//	c - количество готовых соединений.
//
// Пример:
//
//	postgres.NewConnection(
//		"postgres://root@localhost:5432/?sslmode=disable",
//		"app-name",
//		postgres.WithMaxIdleConnections(10),
//	)
func WithMaxIdleConnections(c int) Option {
	return func(cfg *config) {
		cfg.maxIdleConnections = c
	}
}

// WithReadOnly возвращает функцию-замыкание, которая применяет указанную функцию проверки соединения
// к конфигурации подключения к базе данных и проставляет соответствующий флаг.
//
// Параметры:
//
//	ro - значение параметра readOnly.
//
// Пример:
//
//	c := postgres.NewConnection(
//		"postgres://root@localhost:5432/?sslmode=disable",
//		"app-name",
//		postgres.WithReadOnly(true),
//	)
func WithReadOnly(ro bool) Option {
	return func(cfg *config) {
		cfg.readOnly = ro
	}
}

// WithoutValidateConnectFunc возвращает функцию-замыкание, которая устанавливает значение функции проверки
// соединения в nil. Это может быть полезно, если вы хотите отключить проверку соединения.
// Важно: метод следует использовать только с драйвером pgx. Если же выбран драйвер bun/pgdriver, то
// для валидации соединения будет использован query-хук из под-пакета reconnect.
//
// Пример:
//
//	c := postgres.NewConnection(
//		"postgres://root@localhost:5432/?sslmode=disable",
//		"app-name",
//		postgres.WithoutValidateConnectFunc(),
//	)
func WithoutValidateConnectFunc() Option {
	return func(cfg *config) {
		cfg.validateConnFunc = nil
	}
}

// WithQueryExecMode возвращает функцию-замыкание, которая устанавливает значение поля queryExecMode.
// Это может быть полезно, если вы работаете с pgbouncer и надо отключить prepared statement.
// Важно: метод следует использовать только с драйвером pgx.
//
// Пример:
//
//	c := postgres.NewConnection(
//		"postgres://root@localhost:5432/?sslmode=disable",
//		"app-name",
//		postgres.WithQueryExecMode(pgx.QueryExecModeSimpleProtocol),
//	)
func WithQueryExecMode(mode pgx.QueryExecMode) Option {
	return func(cfg *config) {
		cfg.queryExecMode = &mode
	}
}

// WithSlowQueryThreshold возвращает функцию-замыкание, которая устанавливает значение
// поля slowQueryTime. Этот параметр задаёт порог длительности для определения медленных запросов.
//
// Пример:
//
//	threshold := 500 * time.Millisecond
//	c := postgres.NewConnection(
//		"postgres://root@localhost:5432/?sslmode=disable",
//		"app-name",
//		postgres.WithSlowQueryThreshold(threshold),
//	)
func WithSlowQueryThreshold(threshold time.Duration) Option {
	return func(cfg *config) {
		cfg.slowQueryThreshold = threshold
	}
}

// WithSlowQueryLogLevel возвращает функцию-замыкание, которая устанавливает значение
// поля slowQueryLogLevel. Этот параметр задаёт уровень логирования для медленных запросов.
//
// Пример:
//
//	c := postgres.NewConnection(
//		"postgres://root@localhost:5432/?sslmode=disable",
//		"app-name",
//		postgres.WithSlowQueryLogLevel(zerolog.WarnLevel),
//	)
func WithSlowQueryLogLevel(level zerolog.Level) Option {
	return func(cfg *config) {
		cfg.slowQueryLogLevel = level
	}
}

// WithQueryErrorLogLevel возвращает функцию-замыкание, которая устанавливает значение
// поля queryErrorLogLevel. Этот параметр задаёт уровень логирования для запросов,
// выполнение которых завершилось с ошибкой.
//
// Пример:
//
//	c := postgres.NewConnection(
//		"postgres://root@localhost:5432/?sslmode=disable",
//		"app-name",
//		postgres.WithQueryErrorLogLevel(zerolog.Disabled),
//	)
func WithQueryErrorLogLevel(level zerolog.Level) Option {
	return func(cfg *config) {
		cfg.queryErrorLogLevel = level
	}
}

// WithQueryLogLevel возвращает функцию-замыкание, которая устанавливает значение
// поля queryLogLevel. Этот параметр задаёт уровень логирования для общих запросов.
//
// Пример:
//
//	c := postgres.NewConnection(
//		"postgres://root@localhost:5432/?sslmode=disable",
//		"app-name",
//		postgres.WithQueryLogLevel(zerolog.Disabled),
//	)
func WithQueryLogLevel(level zerolog.Level) Option {
	return func(cfg *config) {
		cfg.queryLogLevel = level
	}
}

// WithAdditionalSuffix возвращает функцию-замыкание, которая устанавливает значение
// поля withAdditionalSuffix. Этот параметр влияет на подстановку суффикса к имени базы,
// который позволяет разделить при мониторинге мастер/реплику.
//
// Пример:
//
//	c := postgres.NewConnection(
//		"postgres://root@localhost:5432/?sslmode=disable",
//		"app-name",
//		postgres.WithAdditionalSuffix(true),
//	)
func WithAdditionalSuffix(withSuffix bool) Option {
	return func(cfg *config) {
		cfg.withAdditionalSuffix = withSuffix
	}
}

// WithDriverType задаёт драйвер, с помощью которого создаётся driver.Connector.
//
// Параметры:
//
//	driver - для указания требуемого драйвера используются константы PGXDriver (значение по умолчанию)
//	  и BunPGDriver.
//
// Пример:
//
//	c := postgres.NewConnection(
//		"postgres://root@localhost:5432/?sslmode=disable",
//		"app-name",
//		postgres.WithDriverType(postgres.BunPGDriver),
//	)
func WithDriverType(driver DriverType) Option {
	return func(cfg *config) {
		cfg.driverType = driver
	}
}

// WithTrace возвращает функцию-замыкание, которая устанавливает значение
// поля traceEnabled. Этот параметр включает/отключает трассировку запросов.
//
// Пример:
//
//	c := postgres.NewConnection(
//		"postgres://root@localhost:5432/?sslmode=disable",
//		"app-name",
//		postgres.WithTrace(false),
//	)
func WithTraceEnabled(enabled bool) Option {
	return func(cfg *config) {
		cfg.traceEnabled = enabled
	}
}
