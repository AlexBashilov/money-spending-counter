package build

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"

	"booker/internal/app/apiserver"
)

type Server struct {
	router         *mux.Router
	logger         *logrus.Logger
	Transport      http.RoundTripper
	itemsHandler   *apiserver.ItemsHandler
	expenseHandler *apiserver.ExpenseHandler
}

func NewServer(itemsHandler *apiserver.ItemsHandler, expenseHandler *apiserver.ExpenseHandler) *Server {
	s := &Server{
		router:         mux.NewRouter(),
		logger:         logrus.New(),
		itemsHandler:   itemsHandler,
		expenseHandler: expenseHandler,
	}

	s.configureRouter()

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/book_cost_items/create", s.itemsHandler.HandleItemsCreate()).Methods("POST")
	s.router.HandleFunc("/book_cost_items/get_all", s.itemsHandler.HandleGetItems).Methods("GET")
	s.router.HandleFunc("/book_cost_items/get_only_one/{id:[0-9]+}", s.itemsHandler.HandleGetOnlyOneItem).Methods("GET")
	s.router.HandleFunc("/book_cost_items/delete/{id:[0-9]+}", s.itemsHandler.HandleDeleteItems).Methods("DELETE")
	s.router.HandleFunc("/book_cost_items/update/{id:[0-9]+}", s.itemsHandler.HandleItemsUpdate()).Methods("POST")
	s.router.HandleFunc("/book_daily_expense/create", s.expenseHandler.HandleExpenseCreate()).Methods("POST")
	//s.router.HandleFunc("/book_daily_expense/get_by_id/{id:[0-9]+}", s.handleGetExpenseByItem).Methods("GET")
	//s.router.HandleFunc("/book_daily_expense/get_by_date", s.handleGetExpenseByDate).Methods("GET")
	//s.router.HandleFunc("/book_daily_expense/get_by_date_and_item", s.handleGetExpenseByItemAndDate).Methods("GET")
	//s.router.HandleFunc("/book_daily_expense/get_summ_by_period", s.handleGetExpenseSummByPeriod).Methods("GET")
	//s.router.HandleFunc("/book_cost_items/report", s.handleReport).Methods("GET")
	//s.router.HandleFunc("/book_cost_items/report_by_month", s.handleReportByMonth).Methods("GET")
	s.router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}
