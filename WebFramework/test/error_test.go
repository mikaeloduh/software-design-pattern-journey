package test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestCustomNotFound(t *testing.T) {
	router := framework.NewRouter()

	// 創建自定義的 404 處理器
	customHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusNotFound)
		html := `
<!DOCTYPE html>
<html>
<head>
    <title>404 - Page Not Found</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            text-align: center;
            padding: 50px;
        }
        h1 { color: #e74c3c; }
    </style>
</head>
<body>
    <h1>404 - Page Not Found</h1>
    <p>The page you're looking for doesn't exist.</p>
    <p>Requested URL: ` + r.URL.Path + `</p>
</body>
</html>`
		w.Write([]byte(html))
	})

	// 使用自定義的 404 處理中間件
	router.Use(framework.CustomNotFoundMiddleware(customHandler))

	// 創建測試請求
	req := httptest.NewRequest("GET", "/non-existent", nil)
	w := httptest.NewRecorder()

	// 處理請求
	router.ServeHTTP(w, req)

	// 驗證回應
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
	}

	if contentType := w.Header().Get("Content-Type"); contentType != "text/html" {
		t.Errorf("Expected Content-Type %q, got %q", "text/html", contentType)
	}

	// 確認回應包含預期的 HTML 內容
	body := w.Body.String()
	expectedContent := []string{
		"404 - Page Not Found",
		"The page you're looking for doesn't exist",
		"/non-existent", // 請求的 URL 路徑
	}

	for _, content := range expectedContent {
		if !strings.Contains(body, content) {
			t.Errorf("Expected response to contain %q", content)
		}
	}
}
