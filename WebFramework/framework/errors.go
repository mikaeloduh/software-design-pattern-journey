package framework

import "net/http"

// ErrorType 表示錯誤類型
type ErrorType string

// 預定義的錯誤類型
const (
	ErrorTypeNotFound            ErrorType = "NotFound"
	ErrorTypeMethodNotAllowed    ErrorType = "MethodNotAllowed"
	ErrorTypeBadRequest          ErrorType = "BadRequest"
	ErrorTypeUnauthorized        ErrorType = "Unauthorized"
	ErrorTypeForbidden           ErrorType = "Forbidden"
	ErrorTypeInternalServerError ErrorType = "InternalServerError"
)

// Error 代表一個框架錯誤
type Error struct {
	Type    ErrorType
	Message string
	Err     error
}

func (e *Error) Error() string {
	if e.Message != "" {
		return e.Message
	}
	if e.Err != nil {
		return e.Err.Error()
	}
	return string(e.Type)
}

// NewError 創建一個新的錯誤
func NewError(errType ErrorType, message string, err error) *Error {
	return &Error{
		Type:    errType,
		Message: message,
		Err:     err,
	}
}

// DefaultErrorHandler 提供預設的錯誤處理實現
type DefaultErrorHandler struct{}

func (h *DefaultErrorHandler) CanHandle(err ErrorType) bool {
	return true // 預設處理器可以處理所有類型的錯誤
}

func (h *DefaultErrorHandler) HandleError(err *Error, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	switch err.Type {
	case ErrorTypeNotFound:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 page not found"))
	case ErrorTypeMethodNotAllowed:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 method not allowed"))
	case ErrorTypeBadRequest:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request: " + err.Error()))
	case ErrorTypeUnauthorized:
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 unauthorized: " + err.Error()))
	case ErrorTypeForbidden:
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("403 forbidden: " + err.Error()))
	case ErrorTypeInternalServerError:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 internal server error: " + err.Error()))
	default:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 internal server error"))
	}
}

// ErrorAware 定義了可以處理錯誤的介面
type ErrorAware interface {
	HandleError(err *Error, w http.ResponseWriter, r *http.Request)
}

// ErrorHandler interface
type ErrorHandler interface {
	HandleError(err error, c *Context)
}

// JSONErrorHandler is a sample ErrorHandler returning errors in JSON
type JSONErrorHandler struct{}

func (j *JSONErrorHandler) HandleError(err error, c *Context) {
	if err == nil {
		return
	}
	c.JSON(http.StatusInternalServerError, map[string]string{
		"error": err.Error(),
	})
}
