package http

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

// validationErrors converts validator.ValidationErrors into a field → message map.
// Falls back to a generic error key for non-validation errors.
func validationErrors(err error) map[string]string {
	out := make(map[string]string)
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, fe := range ve {
			out[strings.ToLower(fe.Field())] = msgForTag(fe.Tag())
		}
		return out
	}
	out["error"] = err.Error()
	return out
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "this field is required"
	default:
		return "invalid value"
	}
}
