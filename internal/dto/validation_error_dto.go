package dto

type ValidationErrorResponse struct {
	//Message string `json:"message"`
	Field string `json:"field"`
	Error string `json:"error"`
}
