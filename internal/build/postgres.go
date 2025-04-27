package build

import (
	"booker/pkg/postgres"
	"context"

	"github.com/cockroachdb/errors"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
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
