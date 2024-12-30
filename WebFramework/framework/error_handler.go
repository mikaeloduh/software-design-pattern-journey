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

func DefaultHandleErrorFunc(err error, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	switch e, _ := err.(*errors.Error); {
	case e == errors.ErrorTypeNotFound:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 page not found"))
	case e == errors.ErrorTypeMethodNotAllowed:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 method not allowed"))
	case e == errors.ErrorTypeBadRequest:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request: " + err.Error()))
	case e == errors.ErrorTypeUnauthorized:
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 unauthorized: " + err.Error()))
	case e == errors.ErrorTypeForbidden:
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("403 forbidden: " + err.Error()))
	case e == errors.ErrorTypeInternalServerError:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 internal server error: " + err.Error()))
	default:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 internal server error"))
	}
}
