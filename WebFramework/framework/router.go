package framework

import (
	"net/http"
	"strings"

	"webframework/errors"
)

// Handler 是一個可以返回錯誤的 HTTP 處理器
type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request) error
}

// HandlerFunc 是一個實現了 Handler 接口的函數類型
type HandlerFunc func(http.ResponseWriter, *http.Request) error

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	return f(w, r)
}

// WrapHandler 將標準的 http.Handler 轉換為返回 error 的 Handler
func WrapHandler(h http.Handler) Handler {
	return HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		h.ServeHTTP(w, r)
		return nil
	})
}

type Router struct {
	routes        map[string]map[string]Handler
	middlewares   []Middleware
	errorHandlers []ErrorHandler
}

func NewRouter() *Router {
	r := &Router{
		routes:        make(map[string]map[string]Handler),
		errorHandlers: []ErrorHandler{&DefaultErrorHandler{}}, // 預設錯誤處理器
	}
	return r
}

// RegisterErrorHandler 註冊一個錯誤處理器
func (e *Router) RegisterErrorHandler(handler ErrorHandler) {
	e.errorHandlers = append(e.errorHandlers, handler)
}

// HandleError 處理錯誤
func (e *Router) HandleError(err error, w http.ResponseWriter, r *http.Request) {
	// 從後往前查找，讓後註冊的處理器優先處理
	for i := len(e.errorHandlers) - 1; i >= 0; i-- {
		handler := e.errorHandlers[i]
		if handler.CanHandle(err) {
			handler.HandleError(err, w, r)
			return
		}
	}

	// 如果沒有處理器可以處理，使用預設處理器
	(&DefaultErrorHandler{}).HandleError(err, w, r)
}

// Use adds middleware to the router
func (e *Router) Use(middleware ...Middleware) {
	e.middlewares = append(e.middlewares, middleware...)
}

// Handle registers a new route with a matcher for the URL path and method
func (e *Router) Handle(path string, method string, handler interface{}) {
	// 標準化路徑
	path = strings.Trim(path, "/")
	if path == "" {
		path = "/"
	}
	if _, ok := e.routes[path]; !ok {
		e.routes[path] = make(map[string]Handler)
	}

	var h Handler
	switch handler := handler.(type) {
	case Handler:
		h = handler
	case func(http.ResponseWriter, *http.Request) error:
		h = HandlerFunc(handler)
	case http.Handler:
		h = WrapHandler(handler)
	case func(http.ResponseWriter, *http.Request):
		h = WrapHandler(http.HandlerFunc(handler))
	default:
		panic("unsupported handler type")
	}

	e.routes[path][method] = h
}

// ServeHTTP handles incoming HTTP requests and dispatches them to the registered handlers.
func (e *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	if path == "" {
		path = "/"
	}
	method := r.Method

	// 檢查完整路徑
	if methodHandlers, ok := e.routes[path]; ok {
		if h, ok := methodHandlers[method]; ok {
			handler := e.applyMiddleware(h)
			if err := handler.ServeHTTP(w, r); err != nil {
				e.HandleError(err, w, r)
			}
			return
		}
		// 方法不允許
		e.HandleError(errors.ErrorTypeMethodNotAllowed, w, r)
		return
	}

	// 路徑不存在
	e.HandleError(errors.ErrorTypeNotFound, w, r)
}

func (e *Router) applyMiddleware(handler Handler) Handler {
	h := handler
	for i := len(e.middlewares) - 1; i >= 0; i-- {
		mw := e.middlewares[i]
		h = HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
			var err error
			mw(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				err = h.ServeHTTP(w, r)
			})).ServeHTTP(w, r)
			return err
		})
	}
	return h
}
