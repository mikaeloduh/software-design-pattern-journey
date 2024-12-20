package framework

import (
	"net/http"

	"webframework/errors"
)

// ErrorHandler handles errors
type ErrorHandler interface {
	HandleError(err error, c *Context)
}

// DefaultErrorHandler is an ErrorHandler using pain/text format
type DefaultErrorHandler struct{}

func (s *DefaultErrorHandler) HandleError(err error, c *Context) {
	if err == nil {
		return
	}
	if e, ok := err.(*errors.Error); ok {
		code := e.Code
		if code == 0 {
			code = http.StatusInternalServerError
		}
		c.Status(code)
		c.String(e.Error())
	} else {
		c.Status(http.StatusInternalServerError)
		c.String(err.Error())
	}
}

// JSONErrorHandler is a sample ErrorHandler returning errors in JSON
type JSONErrorHandler struct{}

func (j *JSONErrorHandler) HandleError(err error, c *Context) {
	if err == nil {
		return
	}
	if e, ok := err.(*errors.Error); ok {
		code := e.Code
		if code == 0 {
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
