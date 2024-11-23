package test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"webframework/framework"

	"github.com/stretchr/testify/assert"
)

// TestRouting verifies that the routing is correctly set up and that each route returns the expected response.
func TestRouting(t *testing.T) {
	// Create a new ExactMux
	mux := framework.NewExactMux()

	// Register handlers with exact path and method matching
	mux.Handle("/", http.MethodGet, http.HandlerFunc(homeHandler))
	mux.Handle("/hello", http.MethodGet, http.HandlerFunc(helloHandler))

	// Create a sub-mux for "/user"
	userMux := framework.NewExactMux()
	mux.Router("/user", userMux)
	userMux.Handle("/", http.MethodGet, http.HandlerFunc(getUserHandler))
	userMux.Handle("/", http.MethodPost, http.HandlerFunc(postUserHandler))
	userMux.Handle("/profile", http.MethodGet, http.HandlerFunc(userProfileHandler))

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
		{"GET", "/user/profile", http.StatusOK, "User profile page"},
		{"GET", "/user/settings", http.StatusNotFound, "404 page not found\n"},
		{"GET", "/usersomething", http.StatusNotFound, "404 page not found\n"},
		{"GET", "/userextra", http.StatusNotFound, "404 page not found\n"},
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
