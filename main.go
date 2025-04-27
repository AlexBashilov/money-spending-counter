package main

import (
	"booker/cmd"
	_ "booker/docs"
	"booker/internal/config"
	"context"
	"flag"
	"log"
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

	// itemsHandler, expenseHandler := build.BuildNewItemsHandler()

	// err := trace.NewTracer()
	// if err != nil {
	// 	log.Fatal("init tracer", err)
	// }

	// srv := build.NewServer(itemsHandler, expenseHandler)

	// if err := godotenv.Load(".env"); err != nil {
	// 	log.Print("No .env file found")
	// }

	// log.Println("Booker started")

	// if err := http.ListenAndServe(os.Getenv("SERVICE_ADDRESS"), srv); err != nil {
	// 	panic(err)
	// }

	conf, err := config.Load()
	if err != nil {
		panic(err)
	}

	mainContext := context.Background()
	log.Println("the application is launching")

	exitCode := 0

	err = cmd.Run(mainContext, conf)
	if err != nil {
		log.Fatalln(err)

		exitCode = 1
	}

	os.Exit(exitCode)
}
