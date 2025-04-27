package build

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/uptrace/bun"

	"booker/internal/storage/pgstore"
	"booker/pkg/postgres"
)

func (b *Builder) PgBunClient(readonly bool) (*bun.DB, error) {
	withTrace := b.config.UseTrace() && b.config.App.LogLevel == "trace"

	db, err := postgres.NewConnection(
		b.config.Postgres.DSN(readonly),
		b.config.App.Name,
		postgres.WithQueryExecMode(pgx.QueryExecModeSimpleProtocol),
		postgres.WithQueryErrorLogLevel(zerolog.ErrorLevel),
		postgres.WithSlowQueryLogLevel(zerolog.TraceLevel),
		postgres.WithQueryLogLevel(zerolog.TraceLevel),
		postgres.WithTraceEnabled(withTrace),
		postgres.WithReadOnly(readOnlyEnvAware(readonly, b.config.App.Environment)),
	)
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to postgres")
	}

	b.shutdown.add(func(_ context.Context) error {
		if err = db.Close(); err != nil {
			return errors.Wrap(err, "close db connection")
		}

		return nil
	})

	return db, nil
}

func (b *Builder) BuildBunPgCons() (*bun.DB, *bun.DB, error) {
	if b.pgRWConn == nil {
		rw, err := b.PgBunClient(false)
		if err != nil {
			return nil, nil, err
		}

		b.pgRWConn = rw
	}

	if b.pgROConn == nil {
		ro, err := b.PgBunClient(true)
		if err != nil {
			return nil, nil, err
		}

		b.pgROConn = ro
	}

	return b.pgRWConn, b.pgROConn, nil
}

func (b *Builder) BuildStore(rw, ro bun.IDB) (*pgstore.PGStorage, error) {
	return pgstore.NewStorage(ro, rw), nil
}

func readOnlyEnvAware(readonly bool, env string) bool {
	if !readonly {
		return false
	}

	if strings.HasPrefix(env, "local") || strings.HasPrefix(env, "intgr-test") {
		return false
	}

	return true
}
