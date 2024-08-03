package respond

type ErrorItemsResponse struct {
	Error        string `json:"error"`
	ErrorDetails string `json:"error_details"`
}

type ItemsResponse struct {
	Result  string      `json:"result"`
	Details interface{} `json:"details"`
}
