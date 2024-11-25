package apiserver

import (
	_ "booker/docs" // swagger docs
	respond "booker/internal/app/error"
	"booker/internal/app/usecase"
	"booker/model/apiModels"
	"booker/utils/validator"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel/attribute"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel"
)

type ItemsHandler struct {
	service *usecase.Service
}

func NewItemsHandler(service *usecase.Service) *ItemsHandler {
	return &ItemsHandler{service: service}
}

var validate = validator.InitValidator()

var (
	name   = "items-handlers"
	tracer = otel.GetTracerProvider().Tracer(name)
	logger = otelslog.NewLogger(name)
)

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
		_, span := tracer.Start(r.Context(), "book_cost_items/create")
		defer span.End()
		fmt.Println("span", span)
		log.Info().Msg("create items handler")
		logger.Info("create")

		span.SetAttributes(
			attribute.String("http.url", r.URL.String()),
			attribute.String("http.method", r.Method),
		)

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
		respondWithJSON(w, http.StatusBadRequest, respond.ErrorItemsResponse{
			Error:        "Неверные данные для удаления статьи затрат: статья либо удалена, либо не существует",
			ErrorDetails: err.Error()})
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
		Result:  "success",
		Details: fmt.Sprintf("item %d deleted successfully", itemID)})
}

// HandleItemsUpdate UpdateItems    godoc
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
func (s *ItemsHandler) HandleItemsUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eventID, _ := strconv.Atoi(mux.Vars(r)["id"])

		req := &apiModels.CreateItemsRequest{}

		err := validate.Struct(req)
		if err != nil {
			respondWithJSON(w, http.StatusBadRequest, respond.ErrorItemsResponse{
				Error:        "missing required field",
				ErrorDetails: err.Error()})
			return
		}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			respondWithJSON(w, http.StatusBadRequest, respond.ErrorItemsResponse{
				Error:        err.Error(),
				ErrorDetails: "ошибка декодирования боди запроса"})

			return
		}

		if err := s.service.UpdateItems(r.Context(), req, eventID); err != nil {
			respondWithJSON(w, http.StatusBadRequest, respond.ErrorItemsResponse{
				Error:        "ошибка обновления статьи затрат",
				ErrorDetails: err.Error()})

			return
		}
		respondWithJSON(w, http.StatusOK, respond.ItemsResponse{
			Result:  "success",
			Details: "статья затрат успешно обновлена"})
	}
}

// HandleGetOnlyOneItem GetItemsById    godoc
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
func (s *ItemsHandler) HandleGetOnlyOneItem(w http.ResponseWriter, r *http.Request) {
	itemID, _ := strconv.Atoi(mux.Vars(r)["id"])
	res, err := s.service.GetItemsByID(r.Context(), itemID)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, respond.ItemsResponse{
			Result:  "ошибка при обработке запроса",
			Details: err})
		fmt.Println("Никитос проебал багу =) поиск удаленной статьи.", err)
		return
	}
	respondWithJSON(w, http.StatusOK, respond.ItemsResponse{
		Result:  "успешно",
		Details: res})
}
