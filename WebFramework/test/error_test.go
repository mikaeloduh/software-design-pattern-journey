package test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"webframework/framework"

	"github.com/stretchr/testify/assert"
)

func errorProneHandler(w http.ResponseWriter, r *http.Request) {
	panic("simulating internal server error")
}

func userErrorHandler(w http.ResponseWriter, r *http.Request) {
	panic("simulating internal server error")
}

func TestErrorHandling(t *testing.T) {
	// Create a new Router and use the middleware
	route := framework.NewRouter()
	route.Use(framework.RecoverMiddleware)

	route.Handle("/error", http.MethodGet, http.HandlerFunc(errorProneHandler))

	userRoute := framework.NewRouter()
	userRoute.Handle("/error", http.MethodGet, http.HandlerFunc(userErrorHandler))
	route.Router("/user", userRoute)

	// Start a new test server using the custom Router
	ts := httptest.NewServer(route)
	defer ts.Close()

	tests := []struct {
		method       string
		path         string
		expectedCode int
		expectedBody string
	}{
		{"GET", "/error", http.StatusInternalServerError, "internal server error\n"},
		{"GET", "/user/error", http.StatusInternalServerError, "internal server error\n"},
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
