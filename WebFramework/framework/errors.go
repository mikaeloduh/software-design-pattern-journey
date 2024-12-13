package framework

// ErrorType
type ErrorType string

const (
	ErrorTypeNotFound            ErrorType = "NotFound"
	ErrorTypeMethodNotAllowed    ErrorType = "MethodNotAllowed"
	ErrorTypeBadRequest          ErrorType = "BadRequest"
	ErrorTypeUnauthorized        ErrorType = "Unauthorized"
	ErrorTypeForbidden           ErrorType = "Forbidden"
	ErrorTypeInternalServerError ErrorType = "InternalServerError"
)

// Error
type Error struct {
	Type    ErrorType
	Message string
	Err     error
}

func (e *Error) Error() string {
	if e.Message != "" {
		return e.Message
	}
	if e.Err != nil {
		return e.Err.Error()
	}
	return string(e.Type)
}

// NewError
func NewError(errType ErrorType, message string, err error) *Error {
	return &Error{
		Type:    errType,
		Message: message,
		Err:     err,
	}
}
