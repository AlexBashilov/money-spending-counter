package apiserver

import (
	_ "booker/docs" // swagger docs
	respond "booker/internal/app/error"
	"booker/internal/app/usecase"
	"booker/model/apiModels"
	"encoding/json"
	"net/http"
)

type ItemsHandler struct {
	service *usecase.Service
}

func NewItemsHandler(service *usecase.Service) *ItemsHandler {
	return &ItemsHandler{service: service}
}

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

		//itemExist, _ := store.ItemsRepository().CheckExist(req.ItemName)
		//if itemExist {
		//	respondWithJSON(w, http.StatusBadRequest, respond.ErrorItemsResponse{
		//		Error:        "item exist",
		//		ErrorDetails: fmt.Sprintf("added cost items has ununique name - %s", U.ItemName)})
		//	return
		//}

		//guidExist, _ := s.store.Booker().CheckExist(req.GUID)
		//if guidExist {
		//	respondWithJSON(w, http.StatusNotFound, respond.ErrorItemsResponse{
		//		Error:        "guid exist",
		//		ErrorDetails: "enter unique guid"})
		//	return
		//}

		if err := s.service.CreateItems(*req); err != nil {
			respondWithJSON(w, http.StatusUnprocessableEntity, respond.ErrorItemsResponse{
				Error:        err.Error(),
				ErrorDetails: "invalid request body:required request fields not found"})
			return
		}
		respondWithJSON(w, http.StatusCreated, respond.ItemsResponse{
			//Result:  fmt.Sprintf("item %s created with id - %d", U.ItemName, U.ID),
			//Details:
		})
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
//func (s *server) handleGetItems(w http.ResponseWriter, r *http.Request) {
//
//	res, err := repository.ItemsRepository().GetAllItems()
//	if err != nil {
//		s.error(w, r, http.StatusUnprocessableEntity, err)
//	}
//	respondWithJSON(w, http.StatusOK, respond.ItemsResponse{
//		Result:  "success",
//		Details: res})
//}

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
//func (s *server) handleDeleteItems(w http.ResponseWriter, r *http.Request) {
//	itemID, _ := strconv.Atoi(mux.Vars(r)["id"])
//
//	itemExist, _ := s.store.Booker().CheckExist(itemID)
//	if !itemExist {
//		respondWithJSON(w, http.StatusNotFound, respond.ErrorItemsResponse{
//			Error:        "item not found",
//			ErrorDetails: "item deleted or does not exist"})
//		return
//	}
//
//	if err := repository.ItemsRepository.DeleteItems(itemID); err != nil {
//		respondWithJSON(w, http.StatusUnprocessableEntity, respond.ErrorItemsResponse{
//			Error:        err.Error(),
//			ErrorDetails: "something went wrong"})
//		return
//	}
//	expenseExist, _ := s.store.Booker().CheckExpenseExist(itemID)
//	if expenseExist == true {
//		err := repository.ItemsRepository.AddDeletedAt(itemID)
//		if err != nil {
//			log.Print(err)
//		}
//	}
//
//	respondWithJSON(w, http.StatusOK, respond.ItemsResponse{
//		Result:  "deleted",
//		Details: fmt.Sprintf("item %d deleted successfully", itemID)})
//}

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
