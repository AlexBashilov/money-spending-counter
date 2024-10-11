package main

import (
	_ "booker/docs"
	"booker/internal/app/apiserver"
	"booker/internal/build"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
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
	flag.Parse()

	//var validate = validator.InitValidator()

	config := apiserver.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	itemsHandler := build.BuildNewItemsHandler()
	srv := build.NewServer(itemsHandler)

	fmt.Println("Booker started")

	if err = http.ListenAndServe(config.BindAddr, srv); err != nil {
		panic(err)
	}

}
