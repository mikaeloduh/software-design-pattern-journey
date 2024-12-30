package framework

import (
	"net/http"

	"webframework/errors"
)

// ErrorAwareKey 用於從請求上下文中獲取錯誤處理器
const ErrorAwareKey = "errorAware"

// ErrorHandler 定義了錯誤處理器的接口
type ErrorHandler interface {
	// CanHandle 判斷是否可以處理該類型的錯誤
	CanHandle(err error) bool
	// HandleError 處理特定類型的錯誤
	HandleError(err error, w http.ResponseWriter, r *http.Request)
}

// DefaultErrorHandler 提供預設的錯誤處理實現
type DefaultErrorHandler struct{}

func (h *DefaultErrorHandler) CanHandle(err error) bool {
	return true // 預設處理器可以處理所有類型的錯誤
}

func (h *DefaultErrorHandler) HandleError(err error, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	switch e, _ := err.(*errors.Error); {
	case e == errors.ErrorTypeNotFound:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 page not found"))
	case e == errors.ErrorTypeMethodNotAllowed:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 method not allowed"))
	case e == errors.ErrorTypeBadRequest:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 bad request: " + err.Error()))
	case e == errors.ErrorTypeUnauthorized:
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 unauthorized: " + err.Error()))
	case e == errors.ErrorTypeForbidden:
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("403 forbidden: " + err.Error()))
	case e == errors.ErrorTypeInternalServerError:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 internal server error: " + err.Error()))
	default:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 internal server error"))
	}
}
