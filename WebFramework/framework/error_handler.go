package framework

import (
	"webframework/errors"
)

type ErrorHandlerFunc func(err error, c *Context, next func())

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
