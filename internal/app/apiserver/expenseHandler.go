package apiserver

import (
	"booker/internal/app/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

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

// handleGetExpenseByItem GetExpenseByItem    godoc
//
//	@Summary		Get Expense By Item
//	@Description	Get Expense By Item
//
//	@Param			id	path	string	true	"enter item_id"
//
//	@Produce		application/json
//	@Tags			expense
//	@Success		200	{string}	response.Response{}
//	@Failure		422	{string}	response.Response{}
//
//	@Router			/book_daily_expense/get_by_id/{id} [get]
func (s *server) handleGetExpenseByItem(w http.ResponseWriter, r *http.Request) {
	itemID, _ := strconv.Atoi(mux.Vars(r)["id"])

	res, err := s.store.Booker().GetExpenseByItem(itemID)
	if err != nil {
		s.error(w, r, http.StatusUnprocessableEntity, err)
	}
	respondWithJSON(w, http.StatusOK, res)
}

// handleExpenseByDate handleGetExpenseByDate    godoc
//
//	@Summary		Get Expense By date
//	@Description	Get Expense By date
//
//	@Param			request	body	model.ExpensePeriod	true	"query params"
//
//	@Produce		application/json
//	@Tags			expense
//	@Success		200	{string}	response.Response{}
//	@Failure		422	{string}	response.Response{}
//
//	@Router			/daily_expense/get_by_date [get]
func (s *server) handleGetExpenseByDate(w http.ResponseWriter, r *http.Request) {
	type request struct {
		FromDate time.Time `json:"from_date"`
		ToDate   time.Time `json:"to_date"`
	}
	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, r, http.StatusBadRequest, err)
		return
	}

	formattedTime := &model.ExpensePeriod{
		FromDate: req.FromDate,
		ToDate:   req.ToDate,
	}
	res, err := s.store.Booker().GeExpenseByDate(formattedTime)
	if err != nil {
		s.error(w, r, http.StatusUnprocessableEntity, err)
	}
	respondWithJSON(w, http.StatusOK, res)
}

// handleGetExpenseByItemAndDate GetExpenseByItemAndDate    godoc
//
//	@Summary		Get Expense By Item And Date
//	@Description	Get Expense By Item And Date
//	@Param			request	body	model.ExpensePeriod	true	"query params"

// @Produce	application/json
// @Tags		expense
// @Success	200	{string}	response.Response{}
// @Failure	422	{string}	response.Response{}
//
// @Router		/book_daily_expense/get_by_date_and_item [get]
func (s *server) handleGetExpenseByItemAndDate(w http.ResponseWriter, r *http.Request) {
	type request struct {
		FromDate time.Time `json:"from_date"`
		ToDate   time.Time `json:"to_date"`
		Item     string    `json:"item"`
	}
	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, r, http.StatusBadRequest, err)
		return
	}

	timeAndExpense := &model.ExpensePeriod{
		FromDate: req.FromDate,
		ToDate:   req.ToDate,
		Item:     req.Item,
	}
	res, err := s.store.Booker().GeExpenseByItemAndDate(timeAndExpense)
	if err != nil {
		s.error(w, r, http.StatusUnprocessableEntity, err)
	}
	respondWithJSON(w, http.StatusOK, res)
}

// handleGetExpenseSummByPeriod GetExpenseSummByPeriod    godoc
//
//	@Summary		Get Expense Summ By Period
//	@Description	Get Expense Summ By Period
//	@Param			request	body	model.ExpensePeriod	true	"query params"

// @Produce	application/json
// @Tags		expense
// @Success	200	{string}	response.Response{}
// @Failure	422	{string}	response.Response{}
//
// @Router		/book_daily_expense/get_summ_by_period [get]
func (s *server) handleGetExpenseSummByPeriod(w http.ResponseWriter, r *http.Request) {
	type request struct {
		FromDate time.Time `json:"from_date"`
		ToDate   time.Time `json:"to_date"`
		Item     string    `json:"item"`
	}
	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, r, http.StatusBadRequest, err)
		return
	}

	timeAndExpense := &model.ExpensePeriod{
		FromDate: req.FromDate,
		ToDate:   req.ToDate,
		Item:     req.Item,
	}

	if req.Item != "" {
		res, err := s.store.Booker().GetExpenseSummByPeriodAndItem(timeAndExpense)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}
		respondWithJSON(w, http.StatusOK, res)
	} else {
		res, err := s.store.Booker().GetExpenseSummByPeriod(timeAndExpense)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}
		respondWithJSON(w, http.StatusOK, res)
	}

}
