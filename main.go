package main

import (
	_ "booker/docs"
	"booker/internal/build"
	cfg "booker/internal/config"
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/rs/zerolog"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

// @title				Booker Api
// @version			1.0
// @description		This is an items API
// @termsOfService		http://swagger.io/terms/
// @externalDocs.url	https://swagger.io/resources/open-api/
// @host				localhost:8080
func main() {
	conf, err := cfg.Load()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	itemsHandler, expenseHandler := build.BuildNewItemsHandler(conf)

	srv := build.NewServer(itemsHandler, expenseHandler)

	if err := build.Tracer(ctx, conf); err != nil {
		zerolog.Ctx(ctx).Err(err).Msg("connect to otlp")
	}

	log.Println("the application is launching")

	httpAddress := conf.HTTPAddr()
	if err := http.ListenAndServe(httpAddress, srv); err != nil {
		panic(err)
	}

}
