package framework

import (
	"net/http"
	"strings"

	"webframework/errors"
)

type Router struct {
	routes        map[string]map[string]http.Handler
	subRoutes     map[string]*Router
	middlewares   []Middleware
	errorHandlers []ErrorHandler
	parent        *Router // 指向父路由器
}

func NewRouter() *Router {
	r := &Router{
		routes:        make(map[string]map[string]http.Handler),
		subRoutes:     make(map[string]*Router),
		errorHandlers: []ErrorHandler{&DefaultErrorHandler{}}, // 預設錯誤處理器
	}
	// 添加錯誤處理中間件
	r.Use(ErrorAwareMiddleware(r))
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

	// 如果沒有處理器可以處理，且有父路由器，則交給父路由器處理
	if e.parent != nil {
		e.parent.HandleError(err, w, r)
		return
	}

	// 如果都沒有處理器可以處理，使用預設處理器
	(&DefaultErrorHandler{}).HandleError(err, w, r)
}

// Use adds middleware to the router
func (e *Router) Use(middleware ...Middleware) {
	e.middlewares = append(e.middlewares, middleware...)
}

// Handle registers a new route with a matcher for the URL path and method
func (e *Router) Handle(path string, method string, handler http.Handler) {
	// 標準化路徑
	path = strings.Trim(path, "/")
	if path == "" {
		path = "/"
	}
	if _, ok := e.routes[path]; !ok {
		e.routes[path] = make(map[string]http.Handler)
	}
	e.routes[path][method] = handler
}

// Router registers a new sub-router
func (e *Router) Router(path string, subRouter *Router) {
	// 標準化路徑
	path = strings.Trim(path, "/")
	if path == "" {
		path = "/"
	}
	subRouter.parent = e // 設置父路由器
	e.subRoutes[path] = subRouter
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
			handler.ServeHTTP(w, r)
			return
		}
		// 方法不允許
		e.HandleError(errors.ErrorTypeMethodNotAllowed, w, r)
		return
	}

	// 檢查子路由
	segments := strings.Split(path, "/")
	for i := len(segments); i > 0; i-- {
		prefix := strings.Join(segments[:i], "/")
		if prefix == "" {
			prefix = "/"
		}
		if subMux, ok := e.subRoutes[prefix]; ok {
			remainingPath := ""
			if i < len(segments) {
				remainingPath = strings.Join(segments[i:], "/")
			}
			r2 := r.Clone(r.Context())
			r2.URL.Path = "/" + remainingPath
			subMux.ServeHTTP(w, r2)
			return
		}
	}

	// 路徑不存在
	e.HandleError(errors.ErrorTypeNotFound, w, r)
}

func (e *Router) applyMiddleware(handler http.Handler) http.Handler {
	h := handler
	for i := len(e.middlewares) - 1; i >= 0; i-- {
		h = e.middlewares[i](h)
	}
	return h
}

// GetErrorAware 返回一個實現了 ErrorAware 接口的對象
func (e *Router) GetErrorAware() ErrorAware {
	return e
}
