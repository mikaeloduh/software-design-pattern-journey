package framework

import (
	"net/http"
)

// Context holds request/response, param, and chain info.
type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request

	params map[string]string
	Keys   map[string]interface{}

	// For request handlers
	handlers []HandlerFunc
	index    int

	// No separate errorHandler needed; we rely on router.handleErrorChain
	router *Router
}

// HandlerFunc is a normal request handler in the chain
type HandlerFunc func(ctx *Context)

// Next calls the next handler in the chain
func (c *Context) Next() {
	c.index++
	if c.index < len(c.handlers) {
		c.handlers[c.index](c)
	}
}

// Abort stops remaining request handlers
func (c *Context) Abort() {
	c.index = len(c.handlers)
}

// AbortWithError sets an error and triggers the router's error handler chain
func (c *Context) AbortWithError(err error) {
	c.SetError(err)
	c.Abort()
	if c.router != nil {
		c.router.handleErrorChain(err, c)
	}
}

// SetError sets an error in context's Keys
func (c *Context) SetError(err error) {
	c.Keys["error"] = err
}

// Param gets a path param
func (c *Context) Param(key string) string {
	return c.params[key]
}

// Status sets the response status
func (c *Context) Status(code int) {
	c.ResponseWriter.WriteHeader(code)
}

// String writes plain text
func (c *Context) String(s string) {
	c.ResponseWriter.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if _, err := c.ResponseWriter.Write([]byte(s)); err != nil {
		c.AbortWithError(err)
	}
}

// JSON is optional; you can define it similarly
// func (c *Context) JSON(obj interface{}) { ... }
