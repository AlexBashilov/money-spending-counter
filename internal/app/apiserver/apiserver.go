package apiserver

import (
	"booker/internal/app/store/sqlstore"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func Start(config *Config) error {

	db, err := newDB(config.DataBaseURL)

	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.New(db)

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
