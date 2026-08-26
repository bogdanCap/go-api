package response

import (
	"encoding/json"
	"net/http"
	"github.com/bogdanCap/go-api/internal/application/validation"
	"github.com/bogdanCap/go-api/internal/domain"

	"github.com/go-playground/validator/v10"
)

type ValidationMessageFunc func(err validator.FieldError) string


type ValidationErrorResponse struct {
	//Message string `json:"message"`
	Field string `json:"field"`
	Error string `json:"error"`
}


func WriteValidationError(
	w http.ResponseWriter,
	err error,
	messages domain.ValidationError,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity)

	response := ValidationErrorResponse{}

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			response.Field = fieldError.Field()
			response.Error = validation.Message(
				fieldError,
				messages,
			)

			// return only first validation error
			break
		}
	}

	_ = json.NewEncoder(w).Encode(response)
}
