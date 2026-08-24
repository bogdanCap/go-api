package validation

import (
	"github.com/bogdanCap/go-api/internal/domain"

	"github.com/go-playground/validator/v10"
)

func Message(
	err validator.FieldError,
	messages domain.ValidationError,
) string {
	if fieldMessages, ok := messages[err.Field()]; ok {
		if message, ok := fieldMessages[err.Tag()]; ok {
			return message
		}
	}

	return "invalid value"
}
