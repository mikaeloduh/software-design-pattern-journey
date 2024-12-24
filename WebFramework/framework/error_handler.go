package framework

import (
	"fmt"
	"net/http"

	"webframework/errors"
)

type ErrorHandlerFunc func(err error, c *Context, next func())

func DefaultNotFoundErrorHandler(err error, c *Context, next func()) {
	if e, ok := err.(*errors.Error); ok {
		if e == errors.ErrorTypeNotFound {
			c.Status(e.Code)
			c.String(fmt.Sprintf("Cannot find the path \"%v\"", c.Request.URL.Path))
			return
		}
	}
	next()
}

func DefaultMethodNotAllowedErrorHandler(err error, c *Context, next func()) {
	if e, ok := err.(*errors.Error); ok {
		if e == errors.ErrorTypeMethodNotAllowed {
			c.Status(e.Code)
			c.String(fmt.Sprintf("The method \"%v\" is not allowed on the path \"%v\"", c.Request.Method, c.Request.URL.Path))
			return
		}
	}
	next()
}

// DefaultFallbackHandler is our built-in fallback if no one handles the error.
func DefaultFallbackHandler(err error, c *Context, next func()) {
	// We ignore next(), because fallback is the end of chain.
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

// JSONErrorHandlerFunc: an example that forcibly returns JSON for certain errors
func JSONErrorHandlerFunc(err error, c *Context, next func()) {
	if e, ok := err.(*errors.Error); ok {
		// if code in some range => do JSON response and return
		c.Status(e.Code)
		c.JSON(map[string]string{"error": e.Error()})
		return
	}
	// otherwise call next => pass to next error handler
	next()
}
