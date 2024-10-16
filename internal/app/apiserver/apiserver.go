package apiserver

import (
	"booker/internal/app/store/sqlstore"
	"booker/utils/validator"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // pq lib import
	"net/http"
)

var validate = validator.InitValidator()

// Start server
func Start(config *Config) error {

	db, err := newDB(config.DataBaseURL)

	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.New(db)

	//if err := godotenv.Load(".env"); err != nil {
	//	log.Fatal("No .env file found")
	//}
	//
	//_, err = trace.NewTracer(os.Getenv("SERVICE_NAME"), os.Getenv("OTLP_TRACE_ENDPOINT"))
	//if err != nil {
	//	log.Fatalf("unable to initialize tracer provider due: %v", err)
	//}

	srv := newServer(store)

	fmt.Println("Booker started")

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
