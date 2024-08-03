package apiserver

import (
	_ "booker/docs"
	"booker/internal/app/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// HandleItemsCreate CreateItems		godoc
//
//	@Summary		Create items
//	@Description	Create new items data in Db.
//	@Param			input	body	model.UserCostItems	true	"Create items"
//	@Produce		application/json
//	@Tags			items
//	@Success		201	{string}	response.Response{}
//
//	@Failure		422	{string}	response.Response{}
//	@Failure		400	{string}	response.Response{}
//
//	@Router			/book_cost_items/create [post]
func (s *server) HandleItemsCreate() http.HandlerFunc {
	type Request struct {
		ItemName    string `json:"item_name"`
		Code        int    `json:"code"`
		Description string `json:"description"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &Request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		U := &model.UserCostItems{
			ItemName:    req.ItemName,
			Code:        req.Code,
			Description: req.Description,
		}

		if err := s.store.Booker().CreateItems(U); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusCreated, U)
	}
}

// handleGetItems GetAllItems		godoc
//
//	@Summary		Get all items
//	@Description	Get all items recorded to DB
//	@Produce		application/json
//	@Tags			items
//	@Success		200	{string}	response.Response{}
//	@Failure		422	{string}	response.Response{}
//
//	@Router			/book_cost_items/get_all [get]
func (s *server) handleGetItems(w http.ResponseWriter, r *http.Request) {
	res, err := s.store.Booker().GetAllItems()
	if err != nil {
		s.error(w, r, http.StatusUnprocessableEntity, err)
	}
	respondWithJSON(w, http.StatusOK, res)
}

// handleDeleteItems DeleteItems		godoc
//
//	@Summary		Delete item by id
//	@Description	Delete items data from Db.
//	@Param			id	path	string	true	"ID"
//	@Produce		application/json
//	@Tags			items
//	@Success		200	{string}	response.Response{}
//	@Failure		422	{string}	response.Response{}
//
//	@Router			/book_cost_items/delete/{id} [delete]
func (s *server) handleDeleteItems(w http.ResponseWriter, r *http.Request) {
	itemID, _ := strconv.Atoi(mux.Vars(r)["id"])

	if err := s.store.Booker().DeleteItems(itemID); err != nil {
		s.error(w, r, http.StatusUnprocessableEntity, err)
	}

	if err := s.store.Booker().AddDeletedTime(itemID); err != nil {
		s.error(w, r, http.StatusUnprocessableEntity, err)
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "item deleted"})
}

// handleItemsUpdate UpdateItems		godoc
//
//	@Summary		Update Items
//	@Description	Update items data in Db.
//	@Produce		application/json
//	@Tags			items
//	@Param			input	body		model.UserCostItems	true	"Items struct"
//	@Success		200		{string}	response.Response{}
//
//	@Failure		422		{string}	response.Response{}
//	@Failure		400		{string}	response.Response{}
//
//	@Router			/book_cost_items/update/{id} [post]
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

// handleGetOnlyOneItem GetItemsById		godoc
//
//	@Summary		Get Items By Id
//	@Description	Get Items By Id
//
//	@Param			id	path	string	true	"item id"
//
//	@Produce		application/json
//	@Tags			items
//	@Success		200	{string}	response.Response{}
//	@Failure		422	{string}	response.Response{}
//
//	@Router			/book_cost_items/get_only_one/{id} [get]
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
