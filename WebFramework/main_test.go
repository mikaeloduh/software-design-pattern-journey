package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test the homeHandler function
func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(homeHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code 200")
	assert.Equal(t, "Welcome to the homepage!", rr.Body.String(), "Response body does not match")
}

// Test the helloHandler function
func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code 200")
	assert.Equal(t, "Hello, World!", rr.Body.String(), "Response body does not match")
}

// Test the userHandler function for GET request
func TestUserHandlerGet(t *testing.T) {
	req, err := http.NewRequest("GET", "/user", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code 200")
	assert.Equal(t, "Retrieve user information", rr.Body.String(), "Response body does not match")
}

// Test the userHandler function for POST request
func TestUserHandlerPost(t *testing.T) {
	req, err := http.NewRequest("POST", "/user", strings.NewReader(""))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code 200")
	assert.Equal(t, "Create a new user", rr.Body.String(), "Response body does not match")
}

// Test the userHandler function for unsupported methods (e.g., PUT)
func TestUserHandlerMethodNotAllowed(t *testing.T) {
	req, err := http.NewRequest("PUT", "/user", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code, "Expected status code 405")
	assert.Equal(t, "Unsupported request method\n", rr.Body.String(), "Response body does not match")
}
