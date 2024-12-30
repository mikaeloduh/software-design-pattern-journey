package framework

import (
	"net/http"
)

// ErrorChain 實現責任鏈模式的錯誤處理器
type ErrorChain struct {
	handlers       []ErrorHandler
	defaultHandler *DefaultErrorHandler
}

// NewErrorChain 創建一個新的錯誤處理鏈
func NewErrorChain(handlers ...ErrorHandler) *ErrorChain {
	return &ErrorChain{
		handlers:       handlers,
		defaultHandler: &DefaultErrorHandler{},
	}
}

// HandleError 依序嘗試使用處理器處理錯誤
func (c *ErrorChain) HandleError(err error, w http.ResponseWriter, r *http.Request) {
	for _, handler := range c.handlers {
		if handler.CanHandle(err) {
			handler.HandleError(err, w, r)
			return
		}
	}
	// 如果沒有處理器可以處理，使用預設處理器
	c.defaultHandler.HandleError(err, w, r)
}

// AddHandler 添加新的錯誤處理器到鏈的尾部
func (c *ErrorChain) AddHandler(handler ErrorHandler) {
	c.handlers = append(c.handlers, handler)
}
