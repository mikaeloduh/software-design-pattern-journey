package framework

import (
	"net/http"
)

// Global Error Types and Status Codes Mapping Table
var errorStatusMap = map[ErrorType]int{
	ErrorTypeNotFound:            http.StatusNotFound,            // 404
	ErrorTypeMethodNotAllowed:    http.StatusMethodNotAllowed,    // 405
	ErrorTypeBadRequest:          http.StatusBadRequest,          // 400
	ErrorTypeUnauthorized:        http.StatusUnauthorized,        // 401
	ErrorTypeForbidden:           http.StatusForbidden,           // 403
	ErrorTypeInternalServerError: http.StatusInternalServerError, // 500
}

// RegisterErrorResponse allows developers to register or override the response status code for a given ErrorType.
func RegisterErrorResponse(errType ErrorType, statusCode int) {
	errorStatusMap[errType] = statusCode
}

// ErrorHandler handles errors
type ErrorHandler interface {
	HandleError(err error, c *Context)
}

// JSONErrorHandler is a sample ErrorHandler returning errors in JSON
type JSONErrorHandler struct{}

func (j *JSONErrorHandler) HandleError(err error, c *Context) {
	if err == nil {
		return
	}
	if e, ok := err.(*Error); ok {
		code, exists := errorStatusMap[e.Type]
		if !exists {
			code = http.StatusInternalServerError
		}
		c.Status(code)
		c.JSON(map[string]string{
			"error": e.Error(),
		})
	} else {
		c.Status(http.StatusInternalServerError)
		c.JSON(map[string]string{
			"error": err.Error(),
		})
	}
}

// StringErrorHandler is an ErrorHandler using pain/text format
type StringErrorHandler struct{}

func (s *StringErrorHandler) HandleError(err error, c *Context) {
	if err == nil {
		return
	}
	if e, ok := err.(*Error); ok {
		code, exists := errorStatusMap[e.Type]
		if !exists {
			code = http.StatusInternalServerError
		}
		c.Status(code)
		c.String(e.Error())
	} else {
		c.Status(http.StatusInternalServerError)
		c.String(err.Error())
	}
}
