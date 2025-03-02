package build

import (
	"database/sql"
	"strconv"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"

	cfg "booker/internal/config"
)

func PgsqlConnection(cfg cfg.Config) *bun.DB {
	// pgConf, err := pgx.ParseConfig(cfg.Postgres.DSN())
	// if err != nil {
	// 	return nil, errors.Wrap(err, "cannot parse pg config")
	// }

	dbAddr := cfg.PGAddr()

	dsn := pgdriver.NewConnector(
		pgdriver.WithAddr(dbAddr),
		pgdriver.WithUser(cfg.Postgres.Username),
		pgdriver.WithPassword(cfg.Postgres.Password),
		pgdriver.WithDatabase(cfg.Postgres.Database),
		pgdriver.WithInsecure(true),
	)

	sqlDB := sql.OpenDB(dsn)
	bunDB := bun.NewDB(sqlDB, pgdialect.New())

	debug, err := strconv.ParseBool(cfg.Postgres.Debug)
	if err != nil {
		panic(err)
	}

	bunDB.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true), bundebug.WithEnabled(debug)))

	if err := bunDB.Ping(); err != nil {
		panic(err)
	}

	return bunDB

}
