package postgres

import "errors"

// Ошибки создания коннектора к базе данных.
var (
	ErrEmptyDSN     = errors.New("dsn string can't be empty")
	ErrEmptyAppName = errors.New("app name can't be empty")
)
