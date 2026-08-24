package response

import (
	"encoding/json"
	"net/http"
	"github.com/bogdanCap/go-api/internal/application/validation"
	"github.com/bogdanCap/go-api/internal/domain"
	"github.com/bogdanCap/go-api/internal/dto"

	"github.com/go-playground/validator/v10"
)

type ValidationMessageFunc func(err validator.FieldError) string

func WriteValidationError(
	w http.ResponseWriter,
	err error,
	messages domain.ValidationError,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity)

	response := dto.ValidationErrorResponse{}

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
