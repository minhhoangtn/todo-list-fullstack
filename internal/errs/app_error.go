package errs

const (
	CodeNotFound   = "not_found"
	CodeInternal   = "internal"
	CodeBadRequest = "bad_request"
)

type AppError struct {
	Code     string            // machine-readable identifier e.g. CodeNotFound
	UserMsg  string            // safe message for clients
	Internal error             // raw cause — never serialised to responses
	Metadata map[string]string // sanitized context for structured logging
}

func (e *AppError) Error() string { return e.UserMsg }
func (e *AppError) Unwrap() error { return e.Internal }

func NotFound(internal error) *AppError {
	return &AppError{
		Code:     CodeNotFound,
		UserMsg:  "the requested resource was not found",
		Internal: internal,
	}
}

func Internal(internal error) *AppError {
	return &AppError{
		Code:     CodeInternal,
		UserMsg:  "internal service error",
		Internal: internal,
	}
}
