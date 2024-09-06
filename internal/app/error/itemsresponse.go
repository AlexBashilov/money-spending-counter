package respond

// ErrorItemsResponse struct
type ErrorItemsResponse struct {
	Error        string `json:"error"`
	ErrorDetails string `json:"error_details"`
}

// ItemsResponse struct
type ItemsResponse struct {
	Result  string      `json:"result"`
	Details interface{} `json:"details"`
}
