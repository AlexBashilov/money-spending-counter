package apiserver

import (
	_ "booker/docs"
	"booker/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
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
	s.router.HandleFunc("/cost_items/create_items", s.handleItemsCreate()).Methods("POST")
	s.router.HandleFunc("/cost_items/get_all_items", s.handleGetItems).Methods("GET")
	s.router.HandleFunc("/cost_items/get_only_one_items/{id:[0-9]+}/", s.handleGetOnlyOneItem).Methods("GET")
	s.router.HandleFunc("/cost_items/delete_items/{id:[0-9]+}/", s.handleDeleteItems).Methods("DELETE")
	s.router.HandleFunc("/cost_items/update_items/{id:[0-9]+}/", s.handleItemsUpdate()).Methods("POST")
	s.router.HandleFunc("/daily_expense/create_expense", s.handleExpenseCreate()).Methods("POST")
	s.router.HandleFunc("/daily_expense/get_expense_by_id/{id:[0-9]+}/", s.handleGetExpenseByItem).Methods("GET")
	s.router.HandleFunc("/daily_expense/get_expense_by_date/format", s.handleGetExpenseByDate).Methods("GET")
	s.router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
}
