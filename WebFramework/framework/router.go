package framework

import (
	"net/http"
	"strings"

	"webframework/errors"
)

// node represents a route node
type node struct {
	path     string
	children map[string]*node
	handlers map[string]HandlerFunc // method -> handler
}

// Router represents the router with middleware and a chain of error handlers
type Router struct {
	root *node

	// middlewares for normal requests
	middlewares []HandlerFunc

	// errorHandlers is a chain for handling errors
	errorHandlers []ErrorHandlerFunc
}

// NewRouter creates a new Router instance
func NewRouter() *Router {
	return &Router{
		root: &node{
			children: make(map[string]*node),
			handlers: make(map[string]HandlerFunc),
		},
	}
}

// Use adds normal (request) middlewares
func (r *Router) Use(m ...HandlerFunc) {
	r.middlewares = append(r.middlewares, m...)
}

// UseErrorHandler appends chainable error handlers
func (r *Router) UseErrorHandler(handlers ...ErrorHandlerFunc) {
	r.errorHandlers = append(r.errorHandlers, handlers...)
}

// Handle registers a route with a specific HTTP method
func (r *Router) Handle(method string, path string, handler HandlerFunc) {
	trimmed := strings.Trim(path, "/")
	if trimmed == "" {
		// Root route
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

// handleErrorChain executes the chain of error handlers in order
func (r *Router) handleErrorChain(err error, c *Context) {
	if err == nil {
		return
	}
	if len(r.errorHandlers) == 0 {
		// No error handlers, fallback
		DefaultFallbackHandler(err, c, func(error) {})
		return
	}

	var run func(idx int, currentErr error)
	run = func(idx int, currentErr error) {
		if idx >= len(r.errorHandlers) {
			// Done => fallback
			DefaultFallbackHandler(currentErr, c, func(error) {})
			return
		}
		handler := r.errorHandlers[idx]

		// next function takes a newErr and continues chain
		next := func(newErr error) {
			run(idx+1, newErr)
		}

		// call current handler, passing the current error + next
		handler(currentErr, c, next)
	}

	// start from the first
	run(0, err)
}

// Group creates a route group
func (r *Router) Group(prefix string) *RouteGroup {
	return &RouteGroup{
		prefix: prefix,
		router: r,
	}
}

// ServeHTTP implements http.Handler
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := &Context{
		ResponseWriter: w,
		Request:        req,
		Keys:           make(map[string]interface{}),
		index:          -1,
		router:         r,
	}

	pathSegments := strings.Split(strings.Trim(req.URL.Path, "/"), "/")
	if req.URL.Path == "/" {
		pathSegments = []string{}
	}

	n, params := r.matchNode(pathSegments)
	if n == nil {
		// Not found
		ctx.AbortWithError(errors.ErrorTypeNotFound)
		return
	}

	handler := n.handlers[req.Method]
	if handler == nil {
		// Method not allowed
		ctx.AbortWithError(errors.ErrorTypeMethodNotAllowed)
		return
	}

	ctx.params = params

	// Combine global middlewares + final route handler
	ctx.handlers = append(r.middlewares, handler)
	ctx.Next()
}

// matchNode finds the route node matching the path segments
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
			params[paramKey] = seg
		}
		current = child
	}
	return current, params
}

// RouteGroup for grouping routes under a certain prefix
type RouteGroup struct {
	prefix string
	router *Router
	mws    []HandlerFunc
}

// Use adds group-level middlewares
func (g *RouteGroup) Use(m ...HandlerFunc) {
	g.mws = append(g.mws, m...)
}

// Handle registers a handler under group prefix
func (g *RouteGroup) Handle(method, path string, handler HandlerFunc) {
	fullpath := g.prefix + path
	finalHandler := func(c *Context) {
		// prepend group's middlewares to the chain
		finalChain := append(g.mws, handler)
		c.handlers = append(c.handlers, finalChain...)
		c.Next()
	}
	g.router.Handle(method, fullpath, finalHandler)
}

// GET is a shortcut
func (g *RouteGroup) GET(path string, handler HandlerFunc) {
	g.Handle(http.MethodGet, path, handler)
}
