package contract

type ErrorValidation struct {
	Field   string `json:"field"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Code    string             `json:"code"`
	Message string             `json:"message"`
	TraceID string             `json:"trace_id"`
	Errors  []*ErrorValidation `json:"errors"`
}
