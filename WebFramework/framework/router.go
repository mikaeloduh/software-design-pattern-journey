package framework

import (
	"net/http"
	"strings"
)

type Router struct {
	routes    map[string]map[string]http.Handler // path -> method -> handler
	subRoutes map[string]*Router                 // path segment -> sub-mux
}

func NewRouter() *Router {
	return &Router{
		routes:    make(map[string]map[string]http.Handler),
		subRoutes: make(map[string]*Router),
	}
}

// Handle registers a handler for a specific path and method.
func (e *Router) Handle(path string, method string, handler http.Handler) {
	// Remove leading and trailing slashes for consistent storage
	path = strings.Trim(path, "/")
	if _, exists := e.routes[path]; !exists {
		e.routes[path] = make(map[string]http.Handler)
	}
	e.routes[path][method] = handler
}

// Router registers a sub-mux for a specific path segment.
func (e *Router) Router(path string, subMux *Router) {
	// Remove leading and trailing slashes for consistent storage
	path = strings.Trim(path, "/")
	e.subRoutes[path] = subMux
}

// ServeHTTP handles incoming HTTP requests and dispatches them to the registered handlers.
func (e *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Remove leading and trailing slashes from the request path
	path := strings.Trim(r.URL.Path, "/")
	method := r.Method

	// Check for exact route with method
	if methodHandlers, ok := e.routes[path]; ok {
		if handler, ok := methodHandlers[method]; ok {
			handler.ServeHTTP(w, r)
			return
		} else {
			http.Error(w, "Unsupported request method", http.StatusMethodNotAllowed)
			return
		}
	}

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

	// Not found
	http.NotFound(w, r)
}
