package apiserver

import (
	"booker/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/create_items", s.handleItemsCreate()).Methods("POST")
	s.router.HandleFunc("/get_all_items", s.handleGetItems).Methods("GET")
	s.router.HandleFunc("/get_only_one_items/{id:[0-9]+}/", s.handleGetOnlyOneItem).Methods("GET")
	s.router.HandleFunc("/delete_items/{id:[0-9]+}/", s.handleDeleteItems).Methods("DELETE")
	s.router.HandleFunc("/update_items/{id:[0-9]+}/", s.handleItemsUpdate()).Methods("POST")
}
