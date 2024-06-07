package apiserver

import (
	_ "booker/docs"
	"booker/internal/app/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// CreateOrder Items
//
//	@Summary		handleItemsCreate
//	@Description	handleItemsCreate
//	@Tags			items
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.UserCostItems	true	"item_model"
//	@Success		201		{string}	string
//	@Failure		421		{string}	err
//	@Failure		401		{string}	err
//	@Failure		500		{string}	err
//	@Router			/cost_items/create_items/ [post]
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
	itemID, _ := strconv.Atoi(mux.Vars(r)["id"])

	if err := s.store.Booker().DeleteItems(itemID); err != nil {
		s.error(w, r, http.StatusUnprocessableEntity, err)
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "item deleted"})
}

func (s *server) handleItemsUpdate() http.HandlerFunc {
	type request struct {
		ItemName    string `json:"item_name"`
		Code        int    `json:"code"`
		Description string `json:"description"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		eventID, _ := strconv.Atoi(mux.Vars(r)["id"])

		if err := s.store.Booker().DeleteItems(eventID); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}
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
		s.respond(w, r, http.StatusOK, map[string]string{"result": "item updated"})
	}
}

func (s *server) handleGetOnlyOneItem(w http.ResponseWriter, r *http.Request) {
	itemID, _ := strconv.Atoi(mux.Vars(r)["id"])
	res, err := s.store.Booker().GetOnlyOneItem(itemID)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
	}
	if res == nil {
		s.respond(w, r, http.StatusOK, err)
	}
	respondWithJSON(w, http.StatusOK, res)
}
