package framework

import (
	"net/http"
	"strings"

	"webframework/errors"
)

// Handler is a function that implements the Handler interface
type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request) error
}

// HandlerFunc is a function that implements the Handler interface
type HandlerFunc func(http.ResponseWriter, *http.Request) error

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	return f(w, r)
}

// Convert the standard http.Handler to a Handler that returns an error
func WrapHandler(h http.Handler) Handler {
	return HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		h.ServeHTTP(w, r)
		return nil
	})
}

type Router struct {
	routes        map[string]map[string]Handler
	middlewares   []Middleware
	errorHandlers []ErrorHandlerFunc
}

func NewRouter() *Router {
	r := &Router{
		routes:        make(map[string]map[string]Handler),
		errorHandlers: []ErrorHandlerFunc{},
	}
	// register default error handlers
	r.RegisterErrorHandler(DefaultUnauthorizedErrorHandler)
	r.RegisterErrorHandler(DefaultNotFoundErrorHandler)
	r.RegisterErrorHandler(DefaultMethodNotAllowedErrorHandler)
	return r
}

// RegisterErrorHandler register an error handler
func (e *Router) RegisterErrorHandler(handlerFunc ErrorHandlerFunc) {
	// add at the beginning of the handler chain
	e.errorHandlers = append([]ErrorHandlerFunc{handlerFunc}, e.errorHandlers...)
}

// HandleError handles errors
func (e *Router) HandleError(err error, w http.ResponseWriter, r *http.Request) {
	if len(e.errorHandlers) == 0 {
		// use default error handlers if no error handlers
		e.errorHandlers = []ErrorHandlerFunc{DefaultNotFoundErrorHandler, DefaultMethodNotAllowedErrorHandler}
	}

	var currentHandlerIndex = 0
	var next func(error)
	next = func(err error) {
		if currentHandlerIndex >= len(e.errorHandlers) {
			// use default error handler if no error handlers
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		handler := e.errorHandlers[currentHandlerIndex]
		currentHandlerIndex++
		handler(err, w, r, next)
	}

	next(err)
}

// Use adds middleware to the router
func (e *Router) Use(middleware ...Middleware) {
	e.middlewares = append(e.middlewares, middleware...)
}

// Handle registers a new route with a matcher for the URL path and method
func (e *Router) Handle(path string, method string, handler Handler) {
	// 標準化路徑
	path = strings.Trim(path, "/")
	if path == "" {
		path = "/"
	}
	if _, ok := e.routes[path]; !ok {
		e.routes[path] = make(map[string]Handler)
	}

	e.routes[path][method] = handler
}

// ServeHTTP handles incoming HTTP requests and dispatches them to the registered handlers.
func (e *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	if path == "" {
		path = "/"
	}
	method := r.Method

	// check full path
	if methodHandlers, ok := e.routes[path]; ok {
		if h, ok := methodHandlers[method]; ok {
			handler := e.applyMiddleware(h)
			if err := handler.ServeHTTP(w, r); err != nil {
				e.HandleError(err, w, r)
			}
			return
		}
		// 405
		e.HandleError(errors.ErrorTypeMethodNotAllowed, w, r)
		return
	}

	// 404
	e.HandleError(errors.ErrorTypeNotFound, w, r)
}

func (e *Router) applyMiddleware(handler Handler) Handler {
	h := handler
	for i := len(e.middlewares) - 1; i >= 0; i-- {
		mw := e.middlewares[i]
		currentHandler := h
		h = HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
			var err error
			next := func() {
				err = currentHandler.ServeHTTP(w, r)
			}
			if err := mw(w, r, next); err != nil {
				return err
			}
			return err
		})
	}
	return h
}
