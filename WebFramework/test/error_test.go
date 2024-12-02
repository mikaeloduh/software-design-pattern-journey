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
		defer resp.Body.Close()

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err)

		// Check status code and response body
		assert.Equal(t, tc.expectedCode, resp.StatusCode, "Unexpected status code for path %s", tc.path)
		assert.Equal(t, tc.expectedBody, string(body), "Unexpected response body for path %s", tc.path)
	}
}

func TestNotFoundHandler(t *testing.T) {
	route := framework.NewRouter()
	ts := httptest.NewServer(route)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/nonexistent")
	assert.NoError(t, err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.Equal(t, "text/plain; charset=utf-8", resp.Header.Get("Content-Type"))
	assert.Equal(t, "404 page not found", string(body))
}

func TestMethodNotAllowedHandler(t *testing.T) {
	route := framework.NewRouter()

	// Register a handler for GET method only
	route.Handle("test", http.MethodGet, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}))

	ts := httptest.NewServer(route)
	defer ts.Close()

	// Try to access with POST method
	req, err := http.NewRequest(http.MethodPost, ts.URL+"/test", nil)
	assert.NoError(t, err)

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusMethodNotAllowed, resp.StatusCode)
	assert.Equal(t, "text/plain; charset=utf-8", resp.Header.Get("Content-Type"))
	assert.Equal(t, "405 method not allowed", string(body))
}
