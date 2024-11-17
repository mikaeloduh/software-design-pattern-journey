package main

import (
	"io/ioutil"
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

// TestRouting verifies that the routing is correctly set up and that each route returns the expected response.
func TestRouting(t *testing.T) {
	// Create a new ExactMux
	mux := NewExactMux()

	// Register handlers with exact path matching
	mux.Handle("/", http.HandlerFunc(homeHandler))
	mux.Handle("/hello", http.HandlerFunc(helloHandler))
	mux.Handle("/user", http.HandlerFunc(userHandler))

	// Start a new test server using the custom ExactMux
	ts := httptest.NewServer(mux)
	defer ts.Close()

	// Define test cases
	tests := []struct {
		method       string
		path         string
		expectedCode int
		expectedBody string
	}{
		{"GET", "/", http.StatusOK, "Welcome to the homepage!"},
		{"GET", "/hello", http.StatusOK, "Hello, World!"},
		{"GET", "/user", http.StatusOK, "Retrieve user information"},
		{"POST", "/user", http.StatusOK, "Create a new user"},
		{"PUT", "/user", http.StatusMethodNotAllowed, "Unsupported request method\n"},
		{"GET", "/nonexistent", http.StatusNotFound, "404 page not found\n"},
	}

	for _, tc := range tests {
		// Create a new HTTP request
		req, err := http.NewRequest(tc.method, ts.URL+tc.path, nil)
		assert.NoError(t, err)

		// Send the request
		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)

		// Check the status code
		assert.Equal(t, tc.expectedCode, resp.StatusCode, "Unexpected status code for path %s", tc.path)

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err)
		resp.Body.Close()

		// Check the response body
		assert.Equal(t, tc.expectedBody, string(body), "Unexpected response body for path %s", tc.path)
	}
}
