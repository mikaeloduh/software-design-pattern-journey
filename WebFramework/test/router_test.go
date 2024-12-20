package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"webframework/errors"
	"webframework/framework"
)

// mockHandler and mockMiddleware are for testing purposes
func mockHandler(c *framework.Context) {
	c.String("hello world")
}

func mockJSONHandler(c *framework.Context) {
	c.Status(http.StatusOK)
	c.JSON(map[string]string{"message": "ok"})
}

func dynamicParamHandler(c *framework.Context) {
	id := c.Param("id")
	c.Status(http.StatusOK)
	c.JSON(map[string]string{"id": id})
}

func loggerMiddleware(c *framework.Context) {
	// just a demo middleware that does nothing here
	c.Next()
}

func authMiddleware(c *framework.Context) {
	// Suppose we need an auth token
	token := c.Request.Header.Get("X-Token")
	if token != "secret" {
		c.AbortWithError(errors.NewError(http.StatusUnauthorized, fmt.Errorf("invalid token")))
		return
	}
	c.Next()
}

func TestRouter_StaticRoute(t *testing.T) {
	r := framework.NewRouter()
	r.Handle(http.MethodGet, "/", mockHandler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello world", w.Body.String())
}

func TestRouter_DynamicRoute(t *testing.T) {
	// To ensure dynamic route works, we may need to store param keys in the router.
	// Assume the router has been adjusted to parse params correctly.
	r := framework.NewRouter()
	r.Handle(http.MethodGet, "/users/:id", dynamicParamHandler)

	req := httptest.NewRequest(http.MethodGet, "/users/123", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]string
	err := json.NewDecoder(w.Body).Decode(&resp)
	assert.NoError(t, err)
	assert.Equal(t, "123", resp["id"])
}

func TestRouter_Middleware_Global(t *testing.T) {
	r := framework.NewRouter()
	r.Use(loggerMiddleware) // Just testing middleware chaining works
	r.Handle(http.MethodGet, "/", mockHandler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// If middleware breaks something, we won't get "hello world".
	assert.Equal(t, "hello world", w.Body.String())
}

func TestRouter_GroupMiddleware(t *testing.T) {
	r := framework.NewRouter()

	// Group with auth middleware
	g := r.Group("/auth")
	g.Use(authMiddleware)
	g.GET("/secret", mockJSONHandler)

	// Request without token
	req := httptest.NewRequest(http.MethodGet, "/auth/secret", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, "invalid token", w.Body.String())

	// Request with token
	req = httptest.NewRequest(http.MethodGet, "/auth/secret", nil)
	req.Header.Set("X-Token", "secret")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	body, err := ioutil.ReadAll(w.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(body), "ok")
}

func TestContext_JSON(t *testing.T) {
	r := framework.NewRouter()
	r.Handle(http.MethodGet, "/json", mockJSONHandler)

	req := httptest.NewRequest(http.MethodGet, "/json", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}

func TestContext_String(t *testing.T) {
	r := framework.NewRouter()
	r.Handle(http.MethodGet, "/string", mockHandler)

	req := httptest.NewRequest(http.MethodGet, "/string", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Equal(t, "hello world", w.Body.String())
}

// Fake handler to trigger a custom error
func customErrorHandler(c *framework.Context) {
	c.AbortWithError(errors.NewError(http.StatusPaymentRequired, fmt.Errorf("payment is required")))
}

func TestCustomErrorResponse(t *testing.T) {
	r := framework.NewRouter()
	r.Handle(http.MethodGet, "/pay", customErrorHandler)

	req := httptest.NewRequest(http.MethodGet, "/pay", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 402, w.Code)

	assert.Equal(t, "payment is required", w.Body.String())
}
