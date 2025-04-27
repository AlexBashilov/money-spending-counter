package build

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"

	swagger "booker/pkg/restapi"
	"booker/pkg/tracing"
)

func (b *Builder) registerSwaggerEndpoints(router *mux.Router) {
	router.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(swagger.SwaggerFile())
	})
}

func (b *Builder) RestAPIServer(ctx context.Context) (*http.Server, error) {
	server, err := b.HTTPServer(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "creating http server")
	}

	router := b.httpRouter()
	b.registerSwaggerEndpoints(router)

	apiRouter := router.Name("api").Subrouter()

	if b.config.App.WithTrace {
		apiRouter.Use(otelmux.Middleware(b.config.App.Name))
		apiRouter.Use(tracing.HTTPRequestTrace)
		apiRouter.Use(tracing.NewHTTPResponseTrace(b.config.OTLPTrace.ByteLimit))
	}

	apiRouter.PathPrefix("/").Handler(apiRouter)

	return server, nil
}

func (b *Builder) registerBookerHandlers(ctx context.Context, mux *mux.Router) error {
	pgRw, pgRo, err := b.BuildBunPgCons()
	if err != nil {
		return errors.Wrap(err, "creating http handlers")
	}

	repo, err := b.BuildStore(pgRw, pgRo)
	if err != nil {
		return errors.Wrap(err, "creating http handlers")
	}

	logger := zerolog.Ctx(ctx)
	bookerClient := b.NewWmsClient()
	wmsAdapter := b.NewWmsAdapter(wmsClient, logger)
	wmsProxyService := wp.NewService(repo, wmsAdapter, logger)
	handler := wmsproxy.NewHandler(wmsProxyService, logger)

	mux.Methods(http.MethodGet).
		Path("/special/shipmentOrderStatus.json").
		HandlerFunc(handler.ShipmentOrderStatusHandler)

	mux.Methods(http.MethodGet).
		Path("/special/shipmentOrdersStatus.json").
		HandlerFunc(handler.ShipmentOrdersStatusHandler)

	mux.Methods(http.MethodPost).
		Path("/special/retailShipment.json").
		HandlerFunc(handler.RetailShipmentHandler)

	return nil
}
