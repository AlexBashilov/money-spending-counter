package apiserver

import (
	"booker/internal/app/model"
	"booker/internal/app/store"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
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
	s.router.HandleFunc("/delete_items/{id:[0-9]+}/", s.handleDeleteItems).Methods("DELETE")
}

func (s *server) handleItemsCreate() http.HandlerFunc {
	type request struct {
		ItemName    string `json:"item_name"`
		Code        int    `json:"code"`
		Description string `json:"description"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u := &model.UserCostItems{
			ItemName:    req.ItemName,
			Code:        req.Code,
			Description: req.Description,
		}
		if err := s.store.Booker().CreateItems(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) handleGetItems(w http.ResponseWriter, r *http.Request) {
	res, err := s.store.Booker().GetAllItems()
	if err != nil {
		s.error(w, r, http.StatusUnprocessableEntity, err)
	}
	respondWithJSON(w, http.StatusOK, res)
}

func (s *server) handleDeleteItems(w http.ResponseWriter, r *http.Request) {
	eventID, _ := strconv.Atoi(mux.Vars(r)["id"])

	if err := s.store.Booker().DeleteItems(eventID); err != nil {
		s.error(w, r, http.StatusUnprocessableEntity, err)
	}

	respondWithJSON(w, http.StatusOK, http.Response{})
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
