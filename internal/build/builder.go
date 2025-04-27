package build

import (
	"booker/internal/config"
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
)

type Builder struct {
	config config.Config

	shutdown shutdown

	pgRWConn *bun.DB
	pgROConn *bun.DB

	http struct {
		router *mux.Router
		server *http.Server
	}
}

func New(ctx context.Context, conf config.Config) *Builder {
	b := Builder{config: conf} //nolint:exhaustruct

	err := b.tracer(ctx)
	if err != nil {
		zerolog.Ctx(ctx).Err(err).Msg("connect to otlp")
	}

	return &b
}
