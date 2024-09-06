package apiserver

import (
	_ "booker/docs" // swagger docs
	respond "booker/internal/app/error"
	"booker/internal/app/model"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"net/http"
	"strconv"
)

// HandleItemsCreate CreateItems    godoc
//
//	@Summary		Create items
//	@Description	Create new items data in Db.
//
//	@Param			request	body	model.UserCostItems	true	"Query Params"
//
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
		ItemName    string    `json:"item_name"`
		GUID        uuid.UUID `json:"guid"`
		Description string    `json:"description"`
	}
	return func(w http.ResponseWriter, r *http.Request) {

		req := &Request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {

			respondWithJSON(w, http.StatusBadRequest, respond.ErrorItemsResponse{
				Error:        err.Error(),
				ErrorDetails: "invalid (empty) request body"})
			return
		}
		U := &model.UserCostItems{
			ItemName:    req.ItemName,
			GUID:        req.GUID,
			Description: req.Description,
		}
		itemExist, _ := s.store.Booker().CheckExist(req.ItemName)
		if itemExist {
			respondWithJSON(w, http.StatusBadRequest, respond.ErrorItemsResponse{
				Error:        "item exist",
				ErrorDetails: fmt.Sprintf("added cost items has ununique name - %s", U.ItemName)})
			return
		}

		guidExist, _ := s.store.Booker().CheckExist(req.GUID)
		if guidExist {
			respondWithJSON(w, http.StatusNotFound, respond.ErrorItemsResponse{
				Error:        "guid exist",
				ErrorDetails: "enter unique guid"})
			return
		}
		if err := s.store.Booker().CreateItems(U); err != nil {
			respondWithJSON(w, http.StatusUnprocessableEntity, respond.ErrorItemsResponse{
				Error:        err.Error(),
				ErrorDetails: "invalid request body:required request fields not found"})
			return
		}
		respondWithJSON(w, http.StatusCreated, respond.ItemsResponse{
			Result:  fmt.Sprintf("item %s created with id - %d", U.ItemName, U.ID),
			Details: U})
	}
}

// handleGetItems GetAllItems    godoc
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

	_, span := otel.Tracer("GetItems").Start(r.Context(), "firstspan")
	defer span.End()

	res, err := s.store.Booker().GetAllItems()
	if err != nil {
		s.error(w, r, http.StatusUnprocessableEntity, err)
		span.SetAttributes(attribute.KeyValue{Key: "error", Value: attribute.StringValue(err.Error())})
	}
	respondWithJSON(w, http.StatusOK, respond.ItemsResponse{
		Result:  "success",
		Details: res})
}

// handleDeleteItems DeleteItems    godoc
//
//	@Summary		Delete item by id
//	@Description	Delete items data from Db.
//	@Param			id	path	string	true	"Enter item_id"
//	@Produce		application/json
//	@Tags			items
//	@Success		200	{string}	response.Response{}
//	@Failure		422	{string}	response.Response{}
//
//	@Router			/book_cost_items/delete/{id} [delete]
func (s *server) handleDeleteItems(w http.ResponseWriter, r *http.Request) {
	itemID, _ := strconv.Atoi(mux.Vars(r)["id"])

	itemExist, _ := s.store.Booker().CheckExist(itemID)
	if !itemExist {
		respondWithJSON(w, http.StatusNotFound, respond.ErrorItemsResponse{
			Error:        "item not found",
			ErrorDetails: "item deleted or does not exist"})
		return
	}

	if err := s.store.Booker().DeleteItems(itemID); err != nil {
		respondWithJSON(w, http.StatusUnprocessableEntity, respond.ErrorItemsResponse{
			Error:        err.Error(),
			ErrorDetails: "something went wrong"})
		return
	}
	//expenseExist, _ := s.store.Booker().CheckExpenseExist(itemID)
	//if expenseExist == true {
	//	err := s.store.Booker().AddDeletedAt(itemID)
	//	err != nil{
	//		return err
	//	}
	//}

	respondWithJSON(w, http.StatusOK, respond.ItemsResponse{
		Result:  "deleted",
		Details: fmt.Sprintf("item %d deleted successfully", itemID)})
}

// handleItemsUpdate UpdateItems    godoc
//
//	@Summary		Update Items
//	@Description	Update items data in Db.
//	@Produce		application/json
//	@Tags			items
//	@Param			id		path		string				true	"Enter id"
//
//	@Param			request	body		model.UserCostItems	true	"query params"
//
//	@Success		20		{string}	response.Response{}
//
//	@Failure		422		{string}	response.Response{}
//	@Failure		400		{string}	response.Response{}
//
//	@Router			/book_cost_items/update/{id} [post]
func (s *server) handleItemsUpdate() http.HandlerFunc {
	type request struct {
		ItemName    string    `json:"item_name"`
		GUID        uuid.UUID `json:"guid"`
		Description string    `json:"description"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		eventID, _ := strconv.Atoi(mux.Vars(r)["id"])

		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			respondWithJSON(w, http.StatusBadRequest, respond.ErrorItemsResponse{
				Error:        err.Error(),
				ErrorDetails: "invalid request body"})
			return

		}

		itemExist, _ := s.store.Booker().CheckExist(eventID)
		if !itemExist {
			respondWithJSON(w, http.StatusNotFound, respond.ErrorItemsResponse{
				Error:        "item not found",
				ErrorDetails: "item deleted or does not exist"})
			return
		}

		u := &model.UserCostItems{
			ItemName:    req.ItemName,
			GUID:        req.GUID,
			Description: req.Description,
		}

		if _, err := s.store.Booker().UpdateItems(u, eventID); err != nil {
			respondWithJSON(w, http.StatusBadRequest, respond.ErrorItemsResponse{
				Error:        err.Error(),
				ErrorDetails: "can not update item. contact technical support"})
			return
		}
		respondWithJSON(w, http.StatusOK, respond.ItemsResponse{
			Result:  " success",
			Details: "item updated successfully"})
	}
}

// handleGetOnlyOneItem GetItemsById    godoc
//
//	@Summary		Get Items By Id
//	@Description	Get Items By Id
//
//	@Param			id	path	string	true	"Enter item_id"
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
		respondWithJSON(w, http.StatusBadRequest, respond.ItemsResponse{
			Result:  " not found",
			Details: "item not found, deleted or not exist"})
		return
	}
	respondWithJSON(w, http.StatusOK, respond.ItemsResponse{
		Result:  "success",
		Details: res})
}
