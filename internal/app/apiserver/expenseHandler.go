package apiserver

import (
	"booker/internal/app/model"
	"encoding/json"
	"net/http"
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
		s.respond(w, r, http.StatusCreated, u)
	}
}
