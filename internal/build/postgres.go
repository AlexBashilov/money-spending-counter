package build

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func NewStore() *bun.DB {
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found")
	}
	dsn := pgdriver.NewConnector(
		pgdriver.WithAddr(os.Getenv("DB_HOST")),
		pgdriver.WithUser(os.Getenv("DB_USER")),
		pgdriver.WithPassword(os.Getenv("DB_PASS")),
		pgdriver.WithDatabase(os.Getenv("DB_NAME")),
		pgdriver.WithInsecure(true),
	)

	sqlDB := sql.OpenDB(dsn)
	bunDB := bun.NewDB(sqlDB, pgdialect.New())

	debug, err := strconv.ParseBool(os.Getenv("DB_DEBUG"))
	if err != nil {
		panic(err)
	}

	bunDB.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true), bundebug.WithEnabled(debug)))

	if err := bunDB.Ping(); err != nil {
		panic(err)
	}

	return bunDB

}
