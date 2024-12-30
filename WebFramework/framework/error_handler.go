package framework

import (
	"fmt"
	"net/http"

	"webframework/errors"
)

type ErrorHandlerFunc func(err error, w http.ResponseWriter, r *http.Request, next func(error, http.ResponseWriter, *http.Request))

func DefaultNotFoundErrorHandler(err error, w http.ResponseWriter, r *http.Request, next func(error, http.ResponseWriter, *http.Request)) {
	if e, ok := err.(*errors.Error); ok {
		if e == errors.ErrorTypeNotFound {
			w.WriteHeader(e.Code)
			w.Write([]byte(fmt.Sprintf("Cannot find the path \"%v\"", r.URL.Path)))
			return
		}
	}

	next(err, w, r)
}

func DefaultMethodNotAllowedErrorHandler(err error, w http.ResponseWriter, r *http.Request, next func(error, http.ResponseWriter, *http.Request)) {
	if e, ok := err.(*errors.Error); ok {
		if e == errors.ErrorTypeMethodNotAllowed {
			w.WriteHeader(e.Code)
			w.Write([]byte(fmt.Sprintf("Method \"%v\" is not allowed on path \"%v\"", r.Method, r.URL.Path)))
			return
		}
	}

	next(err, w, r)
}
