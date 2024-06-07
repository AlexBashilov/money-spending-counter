package apiserver

import (
	"booker/internal/app/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (s *server) handleExpenseCreate() http.HandlerFunc {
	type request struct {
		Amount float32 `json:"amount"`
		Item   string  `json:"item"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u := &model.UserExpense{
			Amount: req.Amount,
			Item:   req.Item,
			Date:   time.Now(),
		}
		if err := s.store.Booker().CreateExpense(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		if err := s.store.Booker().UpdateItemID(u.Item); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) handleGetExpenseByItem(w http.ResponseWriter, r *http.Request) {
	itemID, _ := strconv.Atoi(mux.Vars(r)["id"])

	res, err := s.store.Booker().GetExpenseByItem(itemID)
	if err != nil {
		s.error(w, r, http.StatusUnprocessableEntity, err)
	}
	respondWithJSON(w, http.StatusOK, res)
}

func (s *server) handleGetExpenseByDate(w http.ResponseWriter, r *http.Request) {
	param := "from="
	layout := "2006-01-02T15:04:05-07:00"
	length := len(param) + len(layout)

	if s := r.URL.RawQuery; len(s) < length || !strings.HasPrefix(s, param) {
		// unexpected query
	}
	dateTime := r.URL.RawQuery[len(param):length]
	formattedTime, err := time.Parse(layout, dateTime)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, err)
	}

	fmt.Println(formattedTime)
	res, err := s.store.Booker().GeExpenseByDate(formattedTime)
	if err != nil {
		s.error(w, r, http.StatusUnprocessableEntity, err)
	}
	respondWithJSON(w, http.StatusOK, res)
}
