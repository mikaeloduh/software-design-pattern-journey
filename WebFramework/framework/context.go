package framework

import (
	"encoding/json"
	"net/http"
)

// Context holds request-specific values, request/response objects and path parameters.
type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request

	// params stores path parameters like ":id"
	params map[string]string

	// Keys can store arbitrary data during a request's lifetime
	Keys map[string]interface{}

	// handlers store middleware/handler chain
	handlers []HandlerFunc
	index    int

	errorHandler ErrorHandler
}

// HandlerFunc is a function that handles a request in the Context
type HandlerFunc func(c *Context)

// Next executes the next handler in the chain
func (c *Context) Next() {
	c.index++
	if c.index < len(c.handlers) {
		c.handlers[c.index](c)
	}
}

// Abort prevents pending handlers from running
func (c *Context) Abort() {
	c.index = len(c.handlers)
}

// AbortWithError aborts and sets an error to be handled by ErrorHandler
func (c *Context) AbortWithError(err error) {
	c.SetError(err)
	c.Abort()
	if c.errorHandler != nil {
		c.errorHandler.HandleError(err, c)
	}
}

// SetError sets an error in context keys
func (c *Context) SetError(err error) {
	c.Keys["error"] = err
}

// Param returns path parameter by name
func (c *Context) Param(key string) string {
	return c.params[key]
}

// JSON sends a JSON response
func (c *Context) JSON(code int, obj interface{}) {
	c.ResponseWriter.Header().Set("Content-Type", "application/json")
	c.ResponseWriter.WriteHeader(code)
	json.NewEncoder(c.ResponseWriter).Encode(obj)
}

// String sends a plain text response
func (c *Context) String(code int, format string) {
	c.ResponseWriter.Header().Set("Content-Type", "text/plain; charset=utf-8")
	c.ResponseWriter.WriteHeader(code)
	c.ResponseWriter.Write([]byte(format))
}
