package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"webframework/errors"
	"webframework/framework"
)

// FakeDbError is defined in this file, used for simulating a "DB error" in tests
type FakeDbError struct {
	msg string
}

func (f FakeDbError) Error() string {
	return f.msg
}

// MyGenericError is also defined here to simulate a generic error
type MyGenericError struct {
	msg string
}

func (m MyGenericError) Error() string {
	return m.msg
}

// Define an interceptor here to detect FakeDbError (declared in this file)
func MyDbErrorInterceptor(err error, c *framework.Context, next func()) {
	if e, ok := err.(*errors.Error); ok && e.Err != nil {
		// Use the FakeDbError defined in this file for type assertion
		if _, isFakeDbError := e.Err.(FakeDbError); isFakeDbError {
			e.Code = http.StatusInternalServerError
			c.Status(e.Code)
			c.String("Database error occurred!")
			return // Intercept and terminate the chain
		}
	}
	next()
}

func MyDefaultErrorCoder(err error, c *framework.Context, next func()) {
	if e, ok := err.(*errors.Error); ok {
		if e.Code == 0 {
			e.Code = http.StatusInternalServerError
		}
	}
	next()
}

// Final fallback
func FinalFallbackHandler(err error, c *framework.Context, next func()) {
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

func TestRouter_MethodNotAllowed(t *testing.T) {
	r := framework.NewRouter()
	r.Handle(http.MethodGet, "/", mockHandler)

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	assert.Equal(t, http.StatusText(http.StatusMethodNotAllowed), w.Body.String())
}

func TestRouter_Custom_NotFound(t *testing.T) {
	r := framework.NewRouter()
	r.UseErrorHandler(framework.JSONErrorHandlerFunc)
	// no routes added

	req := httptest.NewRequest(http.MethodGet, "/nonexistent", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var resp map[string]string
	err := json.NewDecoder(w.Body).Decode(&resp)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusText(http.StatusNotFound), resp["error"])
}

// Test: ensure that a generic error is not mistaken for a DB error
func TestErrorHandlerChain(t *testing.T) {
	r := framework.NewRouter()

	// Register multiple ErrorHandlerFuncs
	r.UseErrorHandler(
		MyDbErrorInterceptor, // First intercept DB errors
		MyDefaultErrorCoder,  // Change code=0 to 500
		FinalFallbackHandler, // fallback
	)

	// Simulated route #1: produce a DB error
	r.Handle(http.MethodGet, "/db-error", func(ctx *framework.Context) {
		dbErr := FakeDbError{"db connection timeout"}
		ctx.AbortWithError(errors.NewError(0, dbErr))
	})

	// Simulated route #2: generic error
	r.Handle(http.MethodGet, "/generic-error", func(ctx *framework.Context) {
		genericErr := MyGenericError{"some generic error"}
		ctx.AbortWithError(errors.NewError(0, genericErr))
	})

	// Simulated route #3: code=400 specified
	r.Handle(http.MethodGet, "/bad-request", func(ctx *framework.Context) {
		ctx.AbortWithError(errors.NewError(http.StatusBadRequest, nil))
	})

	t.Run("DB error => Interceptor immediately responds 500", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/db-error", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, "Database error occurred!", w.Body.String())
	})

	t.Run("Generic error code=0 => changed to 500 => final response", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/generic-error", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Expect MyDefaultErrorCoder to change code=0 to 500, then FinalFallbackHandler to respond
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, "some generic error", w.Body.String())
	})

	t.Run("code=400 => remains unchanged => final response 400", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/bad-request", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, http.StatusText(http.StatusBadRequest), w.Body.String())
	})
}
