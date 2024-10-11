package build

import (
	"booker/internal/app/apiserver"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"net/http"
)

type Server struct {
	router       *mux.Router
	logger       *logrus.Logger
	Transport    http.RoundTripper
	itemsHandler *apiserver.ItemsHandler
}

func NewServer(itemsHandler *apiserver.ItemsHandler) *Server {
	s := &Server{
		router:       mux.NewRouter(),
		logger:       logrus.New(),
		Transport:    otelhttp.NewTransport(http.DefaultTransport),
		itemsHandler: itemsHandler,
	}

	s.configureRouter()

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/book_cost_items/create", s.itemsHandler.HandleItemsCreate()).Methods("POST")
	//s.router.HandleFunc("/book_cost_items/get_all", s.handleGetItems).Methods("GET")
	//s.router.HandleFunc("/book_cost_items/get_only_one/{id:[0-9]+}", s.handleGetOnlyOneItem).Methods("GET")
	//s.router.HandleFunc("/book_cost_items/delete/{id:[0-9]+}", s.handleDeleteItems).Methods("DELETE")
	//s.router.HandleFunc("/book_cost_items/update/{id:[0-9]+}", s.handleItemsUpdate()).Methods("POST")
	//s.router.HandleFunc("/book_daily_expense/create", s.handleExpenseCreate()).Methods("POST")
	//s.router.HandleFunc("/book_daily_expense/get_by_id/{id:[0-9]+}", s.handleGetExpenseByItem).Methods("GET")
	//s.router.HandleFunc("/book_daily_expense/get_by_date", s.handleGetExpenseByDate).Methods("GET")
	//s.router.HandleFunc("/book_daily_expense/get_by_date_and_item", s.handleGetExpenseByItemAndDate).Methods("GET")
	//s.router.HandleFunc("/book_daily_expense/get_summ_by_period", s.handleGetExpenseSummByPeriod).Methods("GET")
	//s.router.HandleFunc("/book_cost_items/report", s.handleReport).Methods("GET")
	//s.router.HandleFunc("/book_cost_items/report_by_month", s.handleReportByMonth).Methods("GET")
	s.router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}
