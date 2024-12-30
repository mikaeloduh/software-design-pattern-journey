package framework

import (
	"fmt"
	"net/http"

	"webframework/errors"
)

type Middleware func(w http.ResponseWriter, r *http.Request, next func()) error

func authMiddleware(w http.ResponseWriter, r *http.Request, next func()) error {
	// Suppose we need an auth token
	token := r.Header.Get("X-Token")
	if token != "secret" {
		return errors.NewError(http.StatusUnauthorized, fmt.Errorf("invalid token"))
	}
	next()
	return nil
}
