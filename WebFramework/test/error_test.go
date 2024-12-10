package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"webframework/framework"
)

// JSONErrorHandler 自定義的 JSON 格式錯誤處理器
type JSONErrorHandler struct {
	errorTypes []framework.ErrorType
}

func NewJSONErrorHandler(types ...framework.ErrorType) *JSONErrorHandler {
	return &JSONErrorHandler{errorTypes: types}
}

func (h *JSONErrorHandler) CanHandle(err framework.ErrorType) bool {
	for _, t := range h.errorTypes {
		if t == err {
			return true
		}
	}
	return false
}

func (h *JSONErrorHandler) HandleError(err *framework.Error, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{
		"error":   string(err.Type),
		"path":    r.URL.String(), // 使用完整 URL
		"message": err.Error(),
	}

	switch err.Type {
	case framework.ErrorTypeNotFound:
		w.WriteHeader(http.StatusNotFound)
	case framework.ErrorTypeMethodNotAllowed:
		w.WriteHeader(http.StatusMethodNotAllowed)
		response["method"] = r.Method
	case framework.ErrorTypeBadRequest:
		w.WriteHeader(http.StatusBadRequest)
	case framework.ErrorTypeUnauthorized:
		w.WriteHeader(http.StatusUnauthorized)
	case framework.ErrorTypeForbidden:
		w.WriteHeader(http.StatusForbidden)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(response)
}

// 模擬處理用戶請求的處理器
func userHandler(w http.ResponseWriter, r *http.Request) {
	// 從上下文中獲取錯誤處理器
	errorAware := r.Context().Value("errorAware").(framework.ErrorAware)

	// 檢查必要的參數
	userID := r.URL.Query().Get("id")
	if userID == "" {
		errorAware.HandleError(
			framework.NewError(
				framework.ErrorTypeBadRequest,
				"missing required parameter: id",
				nil,
			),
			w, r,
		)
		return
	}

	// 模擬用戶驗證
	if userID == "unauthorized" {
		errorAware.HandleError(
			framework.NewError(
				framework.ErrorTypeUnauthorized,
				"invalid credentials",
				nil,
			),
			w, r,
		)
		return
	}

	// 模擬權限檢查
	if userID == "forbidden" {
		errorAware.HandleError(
			framework.NewError(
				framework.ErrorTypeForbidden,
				"access denied",
				nil,
			),
			w, r,
		)
		return
	}

	// 模擬找不到用戶
	if userID == "nonexistent" {
		errorAware.HandleError(
			framework.NewError(
				framework.ErrorTypeNotFound,
				"user not found",
				nil,
			),
			w, r,
		)
		return
	}

	// 正常回應
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"id":      userID,
		"message": "user found",
	})
}

func TestErrorHandling(t *testing.T) {
	router := framework.NewRouter()

	// 註冊自定義的錯誤處理器
	jsonHandler := NewJSONErrorHandler(
		framework.ErrorTypeNotFound,
		framework.ErrorTypeMethodNotAllowed,
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
			expectedError: string(framework.ErrorTypeNotFound),
		},
		{
			name:           "405 error - method not allowed",
			method:         http.MethodPost,
			path:           "/test",
			expectedCode:   http.StatusMethodNotAllowed,
			expectedError:  string(framework.ErrorTypeMethodNotAllowed),
			expectedMethod: http.MethodPost,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			// 檢查狀態碼
			if w.Code != tt.expectedCode {
				t.Errorf("Expected status code %d, got %d", tt.expectedCode, w.Code)
			}

			// 檢查 Content-Type
			if contentType := w.Header().Get("Content-Type"); contentType != "application/json" {
				t.Errorf("Expected Content-Type %q, got %q", "application/json", contentType)
			}

			// 解析回應
			var response map[string]interface{}
			if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
				t.Fatalf("Failed to decode response: %v", err)
			}

			// 檢查錯誤類型
			if response["error"] != tt.expectedError {
				t.Errorf("Expected error %q, got %q", tt.expectedError, response["error"])
			}

			// 檢查路徑
			if response["path"] != tt.path {
				t.Errorf("Expected path %q, got %q", tt.path, response["path"])
			}

			// 對於 405 錯誤，檢查方法
			if tt.expectedMethod != "" {
				if method, ok := response["method"].(string); !ok || method != tt.expectedMethod {
					t.Errorf("Expected method %q, got %q", tt.expectedMethod, method)
				}
			}
		})
	}
}

// 測試多個錯誤處理器的優先級
func TestErrorHandlerPriority(t *testing.T) {
	router := framework.NewRouter()

	// 創建兩個不同的錯誤處理器
	handler1 := NewJSONErrorHandler(framework.ErrorTypeNotFound)
	handler2 := NewJSONErrorHandler(framework.ErrorTypeNotFound)

	// 註冊處理器（後註冊的優先級更高）
	router.RegisterErrorHandler(handler1)
	router.RegisterErrorHandler(handler2)

	// 發送請求到不存在的路徑
	req := httptest.NewRequest(http.MethodGet, "/non-existent", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	// 驗證回應
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
	}

	// 確認使用了正確的處理器
	var response map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response["error"] != string(framework.ErrorTypeNotFound) {
		t.Errorf("Expected error %q, got %q", framework.ErrorTypeNotFound, response["error"])
	}
}

func TestHandlerErrorHandling(t *testing.T) {
	router := framework.NewRouter()

	// 註冊錯誤處理器
	jsonHandler := NewJSONErrorHandler(
		framework.ErrorTypeBadRequest,
		framework.ErrorTypeUnauthorized,
		framework.ErrorTypeForbidden,
		framework.ErrorTypeNotFound,
	)
	router.RegisterErrorHandler(jsonHandler)

	// 註冊用戶處理器
	router.Handle("/user", http.MethodGet, http.HandlerFunc(userHandler))

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
			expectedError: string(framework.ErrorTypeBadRequest),
		},
		{
			name:          "401 error - unauthorized",
			path:          "/user?id=unauthorized",
			expectedCode:  http.StatusUnauthorized,
			expectedError: string(framework.ErrorTypeUnauthorized),
		},
		{
			name:          "403 error - forbidden",
			path:          "/user?id=forbidden",
			expectedCode:  http.StatusForbidden,
			expectedError: string(framework.ErrorTypeForbidden),
		},
		{
			name:          "404 error - user not found",
			path:          "/user?id=nonexistent",
			expectedCode:  http.StatusNotFound,
			expectedError: string(framework.ErrorTypeNotFound),
		},
		{
			name:          "200 success - valid user",
			path:          "/user?id=123",
			expectedCode:  http.StatusOK,
			expectedError: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			// 檢查狀態碼
			if w.Code != tt.expectedCode {
				t.Errorf("Expected status code %d, got %d", tt.expectedCode, w.Code)
			}

			// 解析回應
			var response map[string]interface{}
			if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
				t.Fatalf("Failed to decode response: %v", err)
			}

			// 對於錯誤情況，檢查錯誤類型和消息
			if tt.expectedError != "" {
				if response["error"] != tt.expectedError {
					t.Errorf("Expected error %q, got %q", tt.expectedError, response["error"])
				}
				if response["path"] != tt.path {
					t.Errorf("Expected path %q, got %q", tt.path, response["path"])
				}
				if msg, ok := response["message"].(string); !ok || msg == "" {
					t.Error("Expected non-empty error message")
				}
			} else {
				// 對於成功情況，檢查用戶數據
				if id, ok := response["id"].(string); !ok || id != "123" {
					t.Errorf("Expected user id %q, got %q", "123", id)
				}
				if msg, ok := response["message"].(string); !ok || msg != "user found" {
					t.Errorf("Expected message %q, got %q", "user found", msg)
				}
			}
		})
	}
}
