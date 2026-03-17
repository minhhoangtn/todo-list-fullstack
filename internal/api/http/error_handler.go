package http

import (
	"errors"
	"minhhoangtn/todo-list-fullstack/internal/errs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func handleAppError(c *gin.Context, err error) {
	var appErr *errs.AppError
	if errors.As(err, &appErr) {
		c.JSON(httpStatusForCode(appErr.Code), errorResponse(appErr.UserMsg))
		return
	}
	c.JSON(http.StatusInternalServerError, errorResponse("an unexpected error occurred"))
}

func errorResponse(err string) gin.H {
	return gin.H{"error": err}
}

func httpStatusForCode(code string) int {
	switch code {
	case errs.CodeNotFound:
		return http.StatusNotFound
	case errs.CodeBadRequest:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

// handleValidationErrors converts validator.ValidationErrors into a field → message map.
// Falls back to a generic error key for non-validation errors.
func handleValidationErrors(err error) map[string]string {
	out := make(map[string]string)
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, fe := range ve {
			out[strings.ToLower(fe.Field())] = msgForTag(fe.Tag())
		}
		return out
	}
	out["error"] = "invalid request body"
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
