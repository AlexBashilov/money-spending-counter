package apiserver

import (
	_ "booker/docs" // swagger docs
	respond "booker/internal/app/error"
	"booker/internal/app/usecase"
	"booker/model/apiModels"
	"booker/utils/validator"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ItemsHandler struct {
	service *usecase.Service
}

func NewItemsHandler(service *usecase.Service) *ItemsHandler {
	return &ItemsHandler{service: service}
}

var validate = validator.InitValidator()

// HandleItemsCreate CreateItems    godoc
//
//	@Summary		Create items
//	@Description	Create new items data in Db.
//
//	@Param			request	body	model.CreateItemsRequest	true	"Query Params"
//
//	@Produce		application/json
//	@Tags			items
//	@Success		201	{string}	response.Response{}
//
//	@Failure		422	{string}	response.Response{}
//	@Failure		400	{string}	response.Response{}
//
//	@Router			/book_cost_items/create [post]
func (s *ItemsHandler) HandleItemsCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &apiModels.CreateItemsRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			respondWithJSON(w, http.StatusBadRequest, respond.ErrorItemsResponse{
				Error:        err.Error(),
				ErrorDetails: "invalid request body"})
			return
		}

		err := validate.Struct(req)
		if err != nil {
			respondWithJSON(w, http.StatusBadRequest, respond.ErrorItemsResponse{
				Error:        "missing required field",
				ErrorDetails: err.Error()})
			return
		}

		if err := s.service.CreateItems(r.Context(), *req); err != nil {
			respondWithJSON(w, http.StatusUnprocessableEntity, respond.ErrorItemsResponse{
				Error:        "invalid request body",
				ErrorDetails: err.Error()})
			return
		}
		respondWithJSON(w, http.StatusCreated, respond.ItemsResponse{
			Result:  fmt.Sprintf("item %s created with id - %s", req.ItemName, req.GUID),
			Details: req,
		})
	}
}

// HandleGetItems GetAllItems    godoc
//
//	@Summary		Get all items
//	@Description	Get all items recorded to DB
//	@Produce		application/json
//	@Tags			items
//	@Success		200	{string}	response.Response{}
//	@Failure		422	{string}	response.Response{}
//
//	@Router			/book_cost_items/get_all [get]
func (s *ItemsHandler) HandleGetItems(w http.ResponseWriter, r *http.Request) {

	response, err := s.service.GetAllItems(r.Context())
	if err != nil {
		respondWithJSON(w, http.StatusUnprocessableEntity, err)
	}
	respondWithJSON(w, http.StatusOK, respond.ItemsResponse{
		Result:  "success",
		Details: response})
}

// HandleDeleteItems DeleteItems    godoc
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
func (s *ItemsHandler) HandleDeleteItems(w http.ResponseWriter, r *http.Request) {
	itemID, _ := strconv.Atoi(mux.Vars(r)["id"])

	if err := s.service.DeleteItems(r.Context(), itemID); err != nil {
		respondWithJSON(w, http.StatusUnprocessableEntity, respond.ErrorItemsResponse{
			Error:        err.Error(),
			ErrorDetails: "something went wrong"})
		return
	}
	//expenseExist, _ := s.store.Booker().CheckExpenseExist(itemID)
	//if expenseExist == true {
	//	err := repository.ItemsRepository.AddDeletedAt(itemID)
	//	if err != nil {
	//		log.Print(err)
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
//	@Param			id		path		string						true	"Enter id"
//
//	@Param			request	body		model.CreateItemsRequest	true	"query params"
//
//	@Success		20		{string}	response.Response{}
//
//	@Failure		422		{string}	response.Response{}
//	@Failure		400		{string}	response.Response{}
//
//	@Router			/book_cost_items/update/{id} [post]
//func (s *server) handleItemsUpdate() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		eventID, _ := strconv.Atoi(mux.Vars(r)["id"])
//
//		req := &apiModels.CreateItemsRequest{}
//
//		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
//			respondWithJSON(w, http.StatusBadRequest, respond.ErrorItemsResponse{
//				Error:        err.Error(),
//				ErrorDetails: "invalid request body"})
//			return
//
//		}
//
//		itemExist, _ := repository.ItemsRepository.CheckExist(eventID)
//		if !itemExist {
//			respondWithJSON(w, http.StatusNotFound, respond.ErrorItemsResponse{
//				Error:        "item not found",
//				ErrorDetails: "item deleted or does not exist"})
//			return
//		}
//
//		u := &apiModels.UserCostItems{
//			ItemName:    req.ItemName,
//			GUID:        req.GUID,
//			Description: req.Description,
//		}
//
//		if _, err := repository.ItemsRepository.UpdateItems(u, eventID); err != nil {
//			respondWithJSON(w, http.StatusBadRequest, respond.ErrorItemsResponse{
//				Error:        err.Error(),
//				ErrorDetails: "can not update item. contact technical support"})
//			return
//		}
//		respondWithJSON(w, http.StatusOK, respond.ItemsResponse{
//			Result:  " success",
//			Details: "item updated successfully"})
//	}
//}

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
//func (s *server) handleGetOnlyOneItem(w http.ResponseWriter, r *http.Request) {
//	itemID, _ := strconv.Atoi(mux.Vars(r)["id"])
//	res, err := repository.ItemsRepository.GetOnlyOneItem(itemID)
//	if err != nil {
//		s.error(w, r, http.StatusInternalServerError, err)
//	}
//	if res == nil {
//		respondWithJSON(w, http.StatusBadRequest, respond.ItemsResponse{
//			Result:  " not found",
//			Details: "item not found, deleted or not exist"})
//		return
//	}
//	respondWithJSON(w, http.StatusOK, respond.ItemsResponse{
//		Result:  "success",
//		Details: res})
//}
