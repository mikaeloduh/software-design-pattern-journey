package framework

import (
	"fmt"
	"net/http"
	"strings"

	"webframework/errors"
)

// ErrorHandlerFunc is an interface of error handler
type ErrorHandlerFunc func(err error, w http.ResponseWriter, r *http.Request, next func(error))

// DefaultNotFoundErrorHandler return 404 page not found with detail message
func DefaultNotFoundErrorHandler(err error, w http.ResponseWriter, r *http.Request, next func(error)) {
	if e, ok := err.(*errors.Error); ok {
		if e == errors.ErrorTypeNotFound {
			w.WriteHeader(e.Code)
			w.Write([]byte(fmt.Sprintf("Cannot find the path \"%v\"", r.URL.Path)))
			return
		}
	}

	next(err)
}

// DefaultMethodNotAllowedErrorHandler return 405 method not allowed with detail message
func DefaultMethodNotAllowedErrorHandler(err error, w http.ResponseWriter, r *http.Request, next func(error)) {
	if e, ok := err.(*errors.Error); ok {
		if e == errors.ErrorTypeMethodNotAllowed {
			w.WriteHeader(e.Code)
			path := strings.Trim(r.URL.Path, "/")
			if path == "" {
				path = "/"
			}
			w.Write([]byte(fmt.Sprintf("Method \"%v\" is not allowed on path \"%v\"", r.Method, path)))
			return
		}
	}

	next(err)
}

func DefaultUnauthorizedErrorHandler(err error, w http.ResponseWriter, r *http.Request, next func(error)) {
	if e, ok := err.(*errors.Error); ok {
		if e == errors.ErrorTypeUnauthorized {
			w.WriteHeader(e.Code)
			w.Write([]byte("401 unauthorized"))
			return
		}
	}

	next(err)
}
