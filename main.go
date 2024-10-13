package main

import (
	_ "booker/docs"
	"booker/internal/build"
	"flag"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
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
	itemsHandler := build.BuildNewItemsHandler()

	srv := build.NewServer(itemsHandler)

	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found")
	}

	log.Println("Booker started")

	if err := http.ListenAndServe(os.Getenv("SERVICE_ADDRESS"), srv); err != nil {
		panic(err)
	}

}
