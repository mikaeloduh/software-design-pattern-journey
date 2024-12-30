package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"webframework/errors"
	"webframework/framework"

	"github.com/stretchr/testify/assert"
)

func TestDefaultErrorHandler(t *testing.T) {
	handler := framework.DefaultHandleErrorFunc
	tests := []struct {
		name           string
		err            error
		expectedCode   int
		expectedStatus string
	}{
		{
			name:           "400 error",
			err:            errors.ErrorTypeBadRequest,
			expectedCode:   http.StatusBadRequest,
			expectedStatus: "400 Bad Request",
		},
		{
			name:           "401 error",
			err:            errors.ErrorTypeUnauthorized,
			expectedCode:   http.StatusUnauthorized,
			expectedStatus: "401 Unauthorized",
		},
		{
			name:           "403 error",
			err:            errors.ErrorTypeForbidden,
			expectedCode:   http.StatusForbidden,
			expectedStatus: "403 Forbidden",
		},
		{
			name:           "404 error",
			err:            errors.ErrorTypeNotFound,
			expectedCode:   http.StatusNotFound,
			expectedStatus: "404 Not Found",
		},
		{
			name:           "500 error",
			err:            errors.ErrorTypeInternalServerError,
			expectedCode:   http.StatusInternalServerError,
			expectedStatus: "500 Internal Server Error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r, err := http.NewRequest("GET", "/", nil)
			assert.NoError(t, err)

			handler(test.err, w, r)

			assert.Equal(t, test.expectedCode, w.Code)
			assert.Equal(t, test.expectedStatus, w.Result().Status)
		})
	}
}

// JSONErrorHandler 自定義的 JSON 格式錯誤處理器
type JSONErrorHandler struct {
	errorTypes []*errors.Error
}

func NewJSONErrorHandler(types ...*errors.Error) *JSONErrorHandler {
	return &JSONErrorHandler{errorTypes: types}
}

func (h *JSONErrorHandler) CanHandle(err error) bool {
	if e, ok := err.(*errors.Error); ok {
		for _, t := range h.errorTypes {
			if t == e {
				return true
			}
		}
	}
	return false
}

func (h *JSONErrorHandler) HandleError(err error, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	e, ok := err.(*errors.Error)
	if !ok {
		// 如果不是預期的錯誤類型，返回一個通用的內部伺服器錯誤
		e = errors.ErrorTypeInternalServerError
	}

	w.WriteHeader(e.Code)

	var message string
	switch e {
	case errors.ErrorTypeNotFound:
		message = "404 page not found"
	case errors.ErrorTypeMethodNotAllowed:
		message = "405 method not allowed"
	case errors.ErrorTypeBadRequest:
		message = "400 bad request"
	case errors.ErrorTypeUnauthorized:
		message = "401 unauthorized"
	case errors.ErrorTypeForbidden:
		message = "403 forbidden"
	case errors.ErrorTypeInternalServerError:
		message = "500 internal server error"
	default:
		message = e.Error()
	}

	response := map[string]interface{}{
		"error":   message,
		"path":    r.URL.String(),
		"message": e.Error(),
	}

	if e == errors.ErrorTypeMethodNotAllowed {
		response["method"] = r.Method
	}

	json.NewEncoder(w).Encode(response)
}

// 模擬處理用戶請求的處理器
func userHandler(w http.ResponseWriter, r *http.Request) error {
	// 從查詢參數中獲取用戶 ID
	userID := r.URL.Query().Get("id")

	// 檢查必要參數
	if userID == "" {
		return errors.ErrorTypeBadRequest
	}

	// 檢查授權
	if userID == "unauthorized" {
		return errors.ErrorTypeUnauthorized
	}

	// 檢查權限
	if userID == "forbidden" {
		return errors.ErrorTypeForbidden
	}

	// 檢查用戶是否存在
	if userID == "nonexistent" {
		return errors.ErrorTypeNotFound
	}

	// 模擬資料庫錯誤
	if userID == "dberror" {
		return errors.ErrorTypeInternalServerError
	}

	// 正常情況下返回用戶資訊
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":      userID,
		"message": "user found",
	})
	return nil
}

func TestCustomErrorHandling(t *testing.T) {
	router := framework.NewRouter()

	// 註冊自定義的錯誤處理器
	jsonHandler := NewJSONErrorHandler(
		errors.ErrorTypeNotFound,
		errors.ErrorTypeMethodNotAllowed,
	)
	router.RegisterErrorHandler(jsonHandler)

	// 註冊一個測試路由
	router.Handle("/test", http.MethodGet, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}))

	tests := []struct {
		name           string
		method         string
		path           string
		expectedCode   int
		expectedError  string
		expectedMethod string
	}{
		{
			name:          "404 error - path not found",
			method:        http.MethodGet,
			path:          "/non-existent",
			expectedCode:  http.StatusNotFound,
			expectedError: "404 page not found",
		},
		{
			name:           "405 error - method not allowed",
			method:         http.MethodPost,
			path:           "/test",
			expectedCode:   http.StatusMethodNotAllowed,
			expectedError:  "405 method not allowed",
			expectedMethod: http.MethodPost,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			// 檢查狀態碼
			assert.Equal(t, tt.expectedCode, w.Code, "Expected status code %d, got %d", tt.expectedCode, w.Code)

			// 檢查 Content-Type
			assert.Equal(t, "application/json", w.Header().Get("Content-Type"), "Expected Content-Type %q, got %q", "application/json", w.Header().Get("Content-Type"))

			// 解析回應
			var response map[string]interface{}
			err := json.NewDecoder(w.Body).Decode(&response)
			assert.NoError(t, err, "Failed to decode response: %v", err)

			// 檢查錯誤類型
			assert.Equal(t, tt.expectedError, response["error"], "Expected error %q, got %q", tt.expectedError, response["error"])

			// 檢查路徑
			assert.Equal(t, tt.path, response["path"], "Expected path %q, got %q", tt.path, response["path"])

			// 對於 405 錯誤，檢查方法
			if tt.expectedMethod != "" {
				method, ok := response["method"].(string)
				assert.True(t, ok, "Expected method to be string")
				assert.Equal(t, tt.expectedMethod, method, "Expected method %q, got %q", tt.expectedMethod, method)
			}
		})
	}
}

// 測試多個錯誤處理器的優先級
func TestErrorHandlerPriority(t *testing.T) {
	router := framework.NewRouter()

	// 創建兩個不同的錯誤處理器
	handler1 := NewJSONErrorHandler(errors.ErrorTypeNotFound)
	handler2 := NewJSONErrorHandler(errors.ErrorTypeNotFound)

	// 註冊處理器（後註冊的優先級更高）
	router.RegisterErrorHandler(handler1)
	router.RegisterErrorHandler(handler2)

	// 發送請求到不存在的路徑
	req := httptest.NewRequest(http.MethodGet, "/non-existent", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	// 驗證回應
	assert.Equal(t, http.StatusNotFound, w.Code, "Expected status code %d, got %d", http.StatusNotFound, w.Code)

	// 確認使用了正確的處理器
	var response map[string]interface{}
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err, "Failed to decode response: %v", err)

	assert.Equal(t, "404 page not found", response["error"], "Expected error %q, got %q", "404 page not found", response["error"])
}

func TestHandlerErrorHandling(t *testing.T) {
	router := framework.NewRouter()

	// 註冊錯誤處理器
	jsonHandler := NewJSONErrorHandler(
		errors.ErrorTypeBadRequest,
		errors.ErrorTypeUnauthorized,
		errors.ErrorTypeForbidden,
		errors.ErrorTypeNotFound,
	)
	router.RegisterErrorHandler(jsonHandler)

	// 註冊用戶處理器
	router.Handle("/user", http.MethodGet, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := userHandler(w, r); err != nil {
			jsonHandler.HandleError(err, w, r)
		}
	}))

	tests := []struct {
		name           string
		path           string
		expectedCode   int
		expectedError  string
		expectedMethod string
	}{
		{
			name:          "400 error - missing id",
			path:          "/user",
			expectedCode:  http.StatusBadRequest,
			expectedError: "400 bad request",
		},
		{
			name:          "401 error - unauthorized",
			path:          "/user?id=unauthorized",
			expectedCode:  http.StatusUnauthorized,
			expectedError: "401 unauthorized",
		},
		{
			name:          "403 error - forbidden",
			path:          "/user?id=forbidden",
			expectedCode:  http.StatusForbidden,
			expectedError: "403 forbidden",
		},
		{
			name:          "404 error - user not found",
			path:          "/user?id=nonexistent",
			expectedCode:  http.StatusNotFound,
			expectedError: "404 page not found",
		},
		{
			name:          "200 success - valid user",
			path:          "/user?id=123",
			expectedCode:  http.StatusOK,
			expectedError: "",
		},
		{
			name:          "500 error - database error",
			path:          "/user?id=dberror",
			expectedCode:  http.StatusInternalServerError,
			expectedError: "500 internal server error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			// 檢查狀態碼
			assert.Equal(t, tt.expectedCode, w.Code, "Expected status code %d, got %d", tt.expectedCode, w.Code)

			// 解析回應
			var response map[string]interface{}
			err := json.NewDecoder(w.Body).Decode(&response)
			assert.NoError(t, err, "Failed to decode response: %v", err)

			// 對於錯誤情況，檢查錯誤類型和消息
			if tt.expectedError != "" {
				assert.Equal(t, tt.expectedError, response["error"], "Expected error %q, got %q", tt.expectedError, response["error"])
				assert.Equal(t, tt.path, response["path"], "Expected path %q, got %q", tt.path, response["path"])
				msg, ok := response["message"].(string)
				assert.True(t, ok && msg != "", "Expected non-empty error message")
			} else {
				// 對於成功情況，檢查用戶數據
				id, ok := response["id"].(string)
				assert.True(t, ok, "Expected id to be string")
				assert.Equal(t, "123", id, "Expected user id %q, got %q", "123", id)

				msg, ok := response["message"].(string)
				assert.True(t, ok, "Expected message to be string")
				assert.Equal(t, "user found", msg, "Expected message %q, got %q", "user found", msg)
			}
		})
	}
}
