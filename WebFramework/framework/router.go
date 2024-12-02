package framework

import (
	"net/http"
	"strings"
)

type Middleware func(http.Handler) http.Handler

type Router struct {
	routes      map[string]map[string]http.Handler // path -> method -> handler
	subRoutes   map[string]*Router                 // path segment -> sub-mux
	middlewares []Middleware
}

func NewRouter() *Router {
	return &Router{
		routes:    make(map[string]map[string]http.Handler),
		subRoutes: make(map[string]*Router),
	}
}

func (e *Router) Use(mw Middleware) {
	e.middlewares = append(e.middlewares, mw)
}

// Handle registers a handler for a specific path and method.
func (e *Router) Handle(path string, method string, handler http.Handler) {
	path = strings.Trim(path, "/")
	if _, exists := e.routes[path]; !exists {
		e.routes[path] = make(map[string]http.Handler)
	}
	e.routes[path][method] = handler
}

// Router registers a sub-mux for a specific path segment.
func (e *Router) Router(path string, subRouter *Router) {
	path = strings.Trim(path, "/")
	// inherent parent middleware
	subRouter.middlewares = append(e.middlewares, subRouter.middlewares...)
	e.subRoutes[path] = subRouter
}

// ServeHTTP handles incoming HTTP requests and dispatches them to the registered handlers.
func (e *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Remove leading and trailing slashes from the request path
	path := strings.Trim(r.URL.Path, "/")
	method := r.Method

	var handler http.Handler

	// Check for exact route with method
	if methodHandlers, ok := e.routes[path]; ok {
		if h, ok := methodHandlers[method]; ok {
			handler = e.applyMiddleware(h)
		} else {
			handler = e.applyMiddleware(http.HandlerFunc(MethodNotAllowedHandler))
		}
	} else {
		// Split the path into segments
		segments := strings.Split(path, "/")

		// Check for sub-mux with matching first path segment
		if len(segments) > 0 {
			firstSegment := segments[0]
			if subMux, ok := e.subRoutes[firstSegment]; ok {
				// Adjust the request URL path
				remainingPath := strings.Join(segments[1:], "/")
				r2 := r.Clone(r.Context())
				r2.URL.Path = "/" + remainingPath
				// Call ServeHTTP on the sub-mux
				subMux.ServeHTTP(w, r2)
				return
			}
		}

		handler = e.applyMiddleware(http.HandlerFunc(NotFoundHandler))
	}

	handler.ServeHTTP(w, r)
}

func (e *Router) applyMiddleware(handler http.Handler) http.Handler {
	for i := len(e.middlewares) - 1; i >= 0; i-- {
		handler = e.middlewares[i](handler)
	}
	return handler
}
