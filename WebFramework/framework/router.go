package framework

import (
	"net/http"
	"strings"
)

// node represents a route node for simple demonstration.
type node struct {
	path     string
	children map[string]*node
	handlers map[string]HandlerFunc // method -> handler
}

// Router represents the router with middleware and error handlers
type Router struct {
	root         *node
	middlewares  []HandlerFunc
	errorHandler ErrorHandler
}

// NewRouter creates a new Router instance
func NewRouter() *Router {
	return &Router{
		root: &node{
			children: make(map[string]*node),
			handlers: make(map[string]HandlerFunc),
		},
		errorHandler: &JSONErrorHandler{},
	}
}

// Use adds global middlewares
func (r *Router) Use(m ...HandlerFunc) {
	r.middlewares = append(r.middlewares, m...)
}

// Handle registers a route with a specific HTTP method
func (r *Router) Handle(method string, path string, handler HandlerFunc) {
	trimmed := strings.Trim(path, "/")
	if trimmed == "" {
		// 是根路由
		r.root.handlers[method] = handler
		return
	}

	segments := strings.Split(trimmed, "/")
	current := r.root
	for _, seg := range segments {
		if current.children[seg] == nil {
			current.children[seg] = &node{
				path:     seg,
				children: make(map[string]*node),
				handlers: make(map[string]HandlerFunc),
			}
		}
		current = current.children[seg]
	}
	current.handlers[method] = handler
}

// Group creates a route group with specific prefix
func (r *Router) Group(prefix string) *RouteGroup {
	return &RouteGroup{
		prefix: prefix,
		router: r,
	}
}

// ServeHTTP implements http.Handler
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := &Context{
		ResponseWriter: w,
		Request:        req,
		Keys:           make(map[string]interface{}),
		errorHandler:   r.errorHandler,
		index:          -1,
	}

	pathSegments := strings.Split(strings.Trim(req.URL.Path, "/"), "/")
	if req.URL.Path == "/" {
		pathSegments = []string{}
	}

	n, params := r.matchNode(pathSegments)
	if n == nil {
		c.AbortWithError(NewError(ErrorTypeNotFound, "not found", nil))
		return
	}

	handler := n.handlers[req.Method]
	if handler == nil {
		c.AbortWithError(NewError(ErrorTypeMethodNotAllowed, "method not supported", nil))
		return
	}

	// 將動態參數放入 c.params
	c.params = params

	// Combine global middlewares and final handler
	c.handlers = append(r.middlewares, handler)
	c.Next()
}
func (r *Router) matchNode(segments []string) (*node, map[string]string) {
	current := r.root
	params := make(map[string]string)

	for _, seg := range segments {
		var child *node
		var paramKey string

		if current.children[seg] != nil {
			child = current.children[seg]
		} else {
			// Check for param node
			for k, ch := range current.children {
				if strings.HasPrefix(k, ":") {
					child = ch
					paramKey = k[1:] // remove ':' prefix
					break
				}
			}
		}

		if child == nil {
			return nil, nil
		}

		if paramKey != "" {
			// Store the actual segment as the param value
			params[paramKey] = seg
		}
		current = child
	}

	return current, params
}

// RouteGroup allows grouping routes
type RouteGroup struct {
	prefix string
	router *Router
	mws    []HandlerFunc
}

// Use adds middleware to the group
func (g *RouteGroup) Use(m ...HandlerFunc) {
	g.mws = append(g.mws, m...)
}

// Handle registers a handler under group prefix
func (g *RouteGroup) Handle(method, path string, handler HandlerFunc) {
	// Combine group prefix + path
	fullpath := g.prefix + path
	// Wrap handler with group's middleware
	finalHandler := func(c *Context) {
		// prepend group's middleware
		finalChain := append(g.mws, handler)
		c.handlers = append(c.handlers, finalChain...)
		c.Next()
	}
	g.router.Handle(method, fullpath, finalHandler)
}

// GET is a shortcut for Handle("GET", path, handler)
func (g *RouteGroup) GET(path string, handler HandlerFunc) {
	g.Handle(http.MethodGet, path, handler)
}
