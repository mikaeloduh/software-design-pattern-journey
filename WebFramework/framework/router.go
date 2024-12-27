package framework

import (
	"net/http"
	"strings"

	"webframework/errors"
)

// Router represents the router with middleware and a chain of error handlers
type Router struct {
	path     string
	children map[string]*Router     // Key: path segment (or :param)
	handlers map[string]HandlerFunc // Key: method (GET, POST...)

	// middlewares for normal requests
	middlewares []HandlerFunc

	// errorHandlers is a chain for handling errors
	errorHandlers []ErrorHandlerFunc
}

// NewRouter creates a new Router instance
func NewRouter() *Router {
	return &Router{
		path:          "",
		children:      make(map[string]*Router),
		handlers:      make(map[string]HandlerFunc),
		middlewares:   make([]HandlerFunc, 0),
		errorHandlers: make([]ErrorHandlerFunc, 0),
	}
}

// Route attaches a sub router at the given path.
func (r *Router) Route(path string, sub *Router) {
	// 先去掉前後 "/"（避免多餘斜線影響 split）
	trimmed := strings.Trim(path, "/")
	if trimmed == "" {
		// 若 path 為空或 "/"
		// 則直接把 sub 視為自己某個 children；依需求可調整
		sub.path = "" // 根路徑
		r.children[sub.path] = sub
		return
	}

	segments := strings.Split(trimmed, "/")

	current := r
	// 一段段往下找/建子路由 (中繼點)，直到最後一段
	for i, seg := range segments {
		if i == len(segments)-1 {
			// 最後一段 => 把 sub 直接掛上去
			sub.path = seg
			current.children[seg] = sub
		} else {
			// 不是最後一段 => 確保中繼 router 存在
			if current.children[seg] == nil {
				current.children[seg] = &Router{
					path:          seg,
					children:      make(map[string]*Router),
					handlers:      make(map[string]HandlerFunc),
					middlewares:   make([]HandlerFunc, 0),
					errorHandlers: make([]ErrorHandlerFunc, 0),
				}
			}
			current = current.children[seg]
		}
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
		r.handlers[method] = handler
		return
	}

	segments := strings.Split(trimmed, "/")

	current := r
	for _, seg := range segments {
		// If child does not exist, create it
		if current.children[seg] == nil {
			current.children[seg] = &Router{
				path:          seg,
				children:      make(map[string]*Router),
				handlers:      make(map[string]HandlerFunc),
				middlewares:   make([]HandlerFunc, 0),
				errorHandlers: make([]ErrorHandlerFunc, 0),
			}
		}
		current = current.children[seg]
	}

	// Finally, assign the handler to the method
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
		router:         r, // root router reference
	}

	chain, params := r.matchRouter(strings.Trim(req.URL.Path, "/"))
	if chain == nil {
		// Not found
		ctx.AbortWithError(errors.ErrorTypeNotFound)
		return
	}

	matchedRouter := chain[len(chain)-1]
	handler := matchedRouter.handlers[req.Method]
	if handler == nil {
		// Method not allowed
		ctx.AbortWithError(errors.ErrorTypeMethodNotAllowed)
		return
	}

	ctx.params = params

	var allMws []HandlerFunc
	for _, routerInChain := range chain {
		allMws = append(allMws, routerInChain.middlewares...)
	}

	ctx.handlers = append(allMws, handler)

	ctx.Next()
}

// matchRouter finds the corresponding subrouter for the given path and any path params
// English comment: We now return a slice of Routers that leads from 'r' (root)
// all the way down to the matched router (if any).
func (r *Router) matchRouter(path string) ([]*Router, map[string]string) {
	if path == "" {
		// means "/"
		return []*Router{r}, make(map[string]string)
	}

	segments := strings.Split(path, "/")
	current := r
	chain := []*Router{current} // keep track of each router we traverse
	params := make(map[string]string)

	for _, seg := range segments {
		if child, ok := current.children[seg]; ok {
			// Exact match
			current = child
			chain = append(chain, current)
		} else {
			// Check if there's a param node
			foundParam := false
			for k, ch := range current.children {
				if strings.HasPrefix(k, ":") {
					paramKey := k[1:]
					params[paramKey] = seg
					current = ch
					chain = append(chain, current)
					foundParam = true
					break
				}
			}
			if !foundParam {
				return nil, nil
			}
		}
	}

	return chain, params
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
