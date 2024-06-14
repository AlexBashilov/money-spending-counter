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
	s.router.HandleFunc("/cost_items/create", s.HandleItemsCreate()).Methods("POST")
	s.router.HandleFunc("/cost_items/get_all", s.handleGetItems).Methods("GET")
	s.router.HandleFunc("/cost_items/get_only_one/{id:[0-9]+}", s.handleGetOnlyOneItem).Methods("GET")
	s.router.HandleFunc("/cost_items/delete/{id:[0-9]+}", s.handleDeleteItems).Methods("DELETE")
	s.router.HandleFunc("/cost_items/update/{id:[0-9]+}", s.handleItemsUpdate()).Methods("POST")
	s.router.HandleFunc("/daily_expense/create", s.handleExpenseCreate()).Methods("POST")
	s.router.HandleFunc("/daily_expense/get_by_id/{id:[0-9]+}", s.handleGetExpenseByItem).Methods("GET")
	s.router.HandleFunc("/daily_expense/get_by_date", s.handleGetExpenseByDate).Methods("GET")
	s.router.HandleFunc("/daily_expense/get_by_date_and_item", s.handleGetExpenseByItemAndDate).Methods("GET")
	s.router.HandleFunc("/daily_expense/get_summ_by_period", s.handleGetExpenseSummByPeriod).Methods("GET")
	s.router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}
