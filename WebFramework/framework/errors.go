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
