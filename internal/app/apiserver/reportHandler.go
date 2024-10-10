package apiserver

import (
	respond "booker/internal/app/error"
	"booker/internal/app/model"
	"encoding/json"
	"net/http"
)

// handleReport report by all expense    godoc
//
//	@Summary		output report
//	@Description	output report by all expense
//	@Produce		application/json
//	@Tags			report
//	@Success		200	{string}	response.Response{}
//	@Failure		422	{string}	response.Response{}
//
//	@Router			/book_cost_items/report [get]
func (s *server) handleReport(w http.ResponseWriter, r *http.Request) {
	res, err := s.store.Booker().GetExpenseSum()
	if err != nil {
		s.error(w, r, http.StatusUnprocessableEntity, err)
	}
	respondWithJSON(w, http.StatusOK, respond.ItemsResponse{
		Result:  "success",
		Details: res})
}

// handleReportByMonth report by expense and month    godoc
//
//	@Summary		report by expense and month
//	@Description	report by expense and month
//	@Produce		application/json
//	@Tags			report
//	@Param			month	path		int	true	"enter month"
//	@Success		200		{string}	response.Response{}
//	@Failure		422		{string}	response.Response{}
//
//	@Router			/book_cost_items/report_by_month [get]
func (s *server) handleReportByMonth(w http.ResponseWriter, r *http.Request) {
	req := &model.ReportByMonth{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		respondWithJSON(w, http.StatusBadRequest, respond.ErrorItemsResponse{
			Error:        err.Error(),
			ErrorDetails: "invalid request body"})
		return
	}

	res, err := s.store.Booker().GetExpenseSumByMonth(req.Month)
	if err != nil {
		s.error(w, r, http.StatusUnprocessableEntity, err)
	}
	respondWithJSON(w, http.StatusOK, respond.ItemsResponse{
		Result:  "success",
		Details: res})
}
