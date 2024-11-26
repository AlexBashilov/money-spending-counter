package apiserver

import (
	respond "booker/internal/app/error"
	"booker/internal/app/usecase"
	"booker/model/apiModels"
	"encoding/json"
	"fmt"
	"net/http"
)

type ExpenseHandler struct {
	service *usecase.Service
}

func NewExpenseHandler(service *usecase.Service) *ExpenseHandler {
	return &ExpenseHandler{service: service}
}

// handleExpenseCreate Expense Create    godoc
//
//	@Summary		Expense Create
//	@Description	Expense Create
//
//	@Param			request	body	model.UserExpense	true	"query params"
//
//	@Produce		application/json
//	@Tags			expense
//	@Success		201	{string}	response.Response{}
//	@Failure		422	{string}	response.Response{}
//	@Failure		400	{string}	response.Response{}
//
//	@Router			/book_daily_expense/create [post]
func (s *ExpenseHandler) HandleExpenseCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &apiModels.CreateExpenseRequest{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			respondWithJSON(w, http.StatusBadRequest, respond.ErrorItemsResponse{
				Error:        "invalid request body",
				ErrorDetails: err.Error()})
			return
		}

		err := validate.Struct(req)
		if err != nil {
			respondWithJSON(w, http.StatusBadRequest, respond.ErrorItemsResponse{
				Error:        "missing required field",
				ErrorDetails: err.Error()})
			return
		}

		if err := s.service.CreateExpense(r.Context(), *req); err != nil {
			respondWithJSON(w, http.StatusUnprocessableEntity, respond.ErrorItemsResponse{
				Error:        "invalid request body",
				ErrorDetails: err.Error()})
			return
		}

		respondWithJSON(w, http.StatusCreated, respond.ItemsResponse{
			Result:  fmt.Sprintf("внесена сумма - %f, по статье - %s", req.Amount, req.Item),
			Details: req,
		})
	}
}

//// handleGetExpenseByItem GetExpenseByItem    godoc
////
////	@Summary		Get Expense By Item
////	@Description	Get Expense By Item
////
////	@Param			id	path	string	true	"enter item_id"
////
////	@Produce		application/json
////	@Tags			expense
////	@Success		200	{string}	response.Response{}
////	@Failure		422	{string}	response.Response{}
////
////	@Router			/book_daily_expense/get_by_id/{id} [get]
//func (s *server) handleGetExpenseByItem(w http.ResponseWriter, r *http.Request) {
//	itemID, _ := strconv.Atoi(mux.Vars(r)["id"])
//
//	res, err := s.store.Booker().GetExpenseByItem(itemID)
//	if err != nil {
//		s.error(w, r, http.StatusUnprocessableEntity, err)
//	}
//	respondWithJSON(w, http.StatusOK, res)
//}
//
//// handleExpenseByDate handleGetExpenseByDate    godoc
////
////	@Summary		Get Expense By date
////	@Description	Get Expense By date
////
////	@Param			request	body	model.ExpensePeriod	true	"query params"
////
////	@Produce		application/json
////	@Tags			expense
////	@Success		200	{string}	response.Response{}
////	@Failure		422	{string}	response.Response{}
////
////	@Router			/daily_expense/get_by_date [get]
//func (s *server) handleGetExpenseByDate(w http.ResponseWriter, r *http.Request) {
//
//	req := &apiModels.GetExpenseByDateRequest{}
//
//	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
//		respondWithJSON(w, http.StatusBadRequest, respond.ErrorItemsResponse{
//			Error:        "invalid request body",
//			ErrorDetails: err.Error()})
//		return
//	}
//
//	err := validate.Struct(req)
//	if err != nil {
//		respondWithJSON(w, http.StatusBadRequest, respond.ErrorItemsResponse{
//			Error:        "missing required field",
//			ErrorDetails: err.Error()})
//		return
//	}
//
//	formattedTime := &apiModels.ExpensePeriod{
//		FromDate: req.FromDate,
//		ToDate:   req.ToDate,
//	}
//	res, err := s.store.Booker().GetExpenseByDate(formattedTime)
//	if err != nil {
//		s.error(w, r, http.StatusUnprocessableEntity, err)
//	}
//	respondWithJSON(w, http.StatusOK, res)
//}
//
//// handleGetExpenseByItemAndDate GetExpenseByItemAndDate    godoc
////
////	@Summary		Get Expense By Item And Date
////	@Description	Get Expense By Item And Date
////	@Param			request	body	model.ExpensePeriod	true	"query params"
//
//// @Produce	application/json
//// @Tags		expense
//// @Success	200	{string}	response.Response{}
//// @Failure	422	{string}	response.Response{}
////
//// @Router		/book_daily_expense/get_by_date_and_item [get]
//func (s *server) handleGetExpenseByItemAndDate(w http.ResponseWriter, r *http.Request) {
//	req := &apiModels.ExpenseItemDateRequest{}
//	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
//		s.error(w, r, http.StatusBadRequest, err)
//		return
//	}
//
//	timeAndExpense := &apiModels.ExpensePeriod{
//		FromDate: req.FromDate,
//		ToDate:   req.ToDate,
//		Item:     req.Item,
//	}
//	res, err := s.store.Booker().GetExpenseByItemAndDate(timeAndExpense)
//	if err != nil {
//		s.error(w, r, http.StatusUnprocessableEntity, err)
//	}
//	respondWithJSON(w, http.StatusOK, res)
//}
//
//// handleGetExpenseSummByPeriod GetExpenseSummByPeriod    godoc
////
////	@Summary		Get Expense Summ By Period
////	@Description	Get Expense Summ By Period
////	@Param			request	body	model.ExpensePeriod	true	"query params"
//
//// @Produce	application/json
//// @Tags		expense
//// @Success	200	{string}	response.Response{}
//// @Failure	422	{string}	response.Response{}
////
//// @Router		/book_daily_expense/get_summ_by_period [get]
//func (s *server) handleGetExpenseSummByPeriod(w http.ResponseWriter, r *http.Request) {
//	req := &apiModels.ExpenseItemDateRequest{}
//	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
//		s.error(w, r, http.StatusBadRequest, err)
//		return
//	}
//
//	timeAndExpense := &apiModels.ExpensePeriod{
//		FromDate: req.FromDate,
//		ToDate:   req.ToDate,
//		Item:     req.Item,
//	}
//
//	if req.Item != "" {
//		res, err := s.store.Booker().GetExpenseSummByPeriodAndItem(timeAndExpense)
//		if err != nil {
//			s.error(w, r, http.StatusUnprocessableEntity, err)
//		}
//		respondWithJSON(w, http.StatusOK, res)
//	} else {
//		res, err := s.store.Booker().GetExpenseSummByPeriod(timeAndExpense)
//		if err != nil {
//			s.error(w, r, http.StatusUnprocessableEntity, err)
//		}
//		respondWithJSON(w, http.StatusOK, res)
//	}
//
//}
