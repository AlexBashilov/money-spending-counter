package main

import (
	_ "booker/docs"
	"booker/internal/build"
	cfg "booker/internal/config"
	"flag"
	"log"
	"net/http"
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

	itemsHandler, expenseHandler := build.BuildNewItemsHandler(conf)

	srv := build.NewServer(itemsHandler, expenseHandler)

	log.Println("Booker started")

	httpAddress := conf.HTTPAddr()
	if err := http.ListenAndServe(httpAddress, srv); err != nil {
		panic(err)
	}

}
