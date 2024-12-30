package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"webframework/framework"

	"github.com/stretchr/testify/assert"
)

// Handler functions remain the same
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage!")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Retrieve user information")
}

func postUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create a new user")
}

func userProfileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User profile page")
}

// TestRouting verifies that the routing is correctly set up and that each route returns the expected response.
func TestRouting(t *testing.T) {
	// Create a new Router
	route := framework.NewRouter()

	// Register handlers with exact path and method matching
	route.Handle("/", http.MethodGet, http.HandlerFunc(homeHandler))
	route.Handle("/hello", http.MethodGet, http.HandlerFunc(helloHandler))
	route.Handle("/user", http.MethodGet, http.HandlerFunc(getUserHandler))
	route.Handle("/user", http.MethodPost, http.HandlerFunc(postUserHandler))
	route.Handle("/user/profile", http.MethodGet, http.HandlerFunc(userProfileHandler))

	// Start a new test server using the custom Router
	ts := httptest.NewServer(route)
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
		{"PUT", "/user", http.StatusMethodNotAllowed, "405 method not allowed"},
		{"GET", "/user/profile", http.StatusOK, "User profile page"},
		{"GET", "/usersomething", http.StatusNotFound, "404 page not found"},
		{"GET", "/userextra", http.StatusNotFound, "404 page not found"},
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
