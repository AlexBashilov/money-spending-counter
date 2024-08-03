package respond

import "booker/internal/app/model"

type ErrorItemsResponse struct {
	Error        string `json:"error"`
	ErrorDetails string `json:"error_details"`
}

type ItemsResponse struct {
	Result  string               `json:"result"`
	Details *model.UserCostItems `json:"details"`
}
