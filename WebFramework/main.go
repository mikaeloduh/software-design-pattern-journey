package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// ExactMux is a custom HTTP request multiplexer that matches exact paths and methods.
type ExactMux struct {
	routes   map[string]map[string]http.Handler // path -> method -> handler
	subMuxes map[string]*ExactMux               // path -> sub-mux
}

// NewExactMux creates a new ExactMux instance.
func NewExactMux() *ExactMux {
	return &ExactMux{
		routes:   make(map[string]map[string]http.Handler),
		subMuxes: make(map[string]*ExactMux),
	}
}

// Handle registers a new route.
// If called with path and handler (*ExactMux), it registers a sub-mux for that path.
// If called with path, method, and handler (http.Handler), it registers a handler for a specific method.
func (e *ExactMux) Handle(path string, args ...interface{}) {
	if len(args) == 1 {
		// Register sub-mux
		if subMux, ok := args[0].(*ExactMux); ok {
			e.subMuxes[path] = subMux
		} else {
			panic("Handle: when providing one argument, it must be of type *ExactMux")
		}
	} else if len(args) == 2 {
		// Register handler for method
		method, ok := args[0].(string)
		if !ok {
			panic("Handle: method must be a string")
		}
		handler, ok := args[1].(http.Handler)
		if !ok {
			panic("Handle: handler must be http.Handler")
		}
		if _, exists := e.routes[path]; !exists {
			e.routes[path] = make(map[string]http.Handler)
		}
		e.routes[path][method] = handler
	} else {
		panic("Handle: incorrect number of arguments")
	}
}

// ServeHTTP handles incoming HTTP requests and dispatches them to the registered handlers.
func (e *ExactMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
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

	// Check for sub-mux
	for subPath, subMux := range e.subMuxes {
		if strings.HasPrefix(path, subPath) {
			// Adjust the request URL path
			r2 := new(http.Request)
			*r2 = *r
			r2.URL = new(url.URL)
			*r2.URL = *r.URL
			r2.URL.Path = strings.TrimPrefix(r.URL.Path, subPath)
			if r2.URL.Path == "" {
				r2.URL.Path = "/"
			}
			subMux.ServeHTTP(w, r2)
			return
		}
	}

	// Not found
	http.NotFound(w, r)
}

func main() {
	// Create a new ExactMux
	mux := NewExactMux()

	// Register handlers with exact path and method matching
	mux.Handle("/", http.MethodGet, http.HandlerFunc(homeHandler))
	mux.Handle("/hello", http.MethodGet, http.HandlerFunc(helloHandler))

	// Create a sub-mux for "/user"
	userMux := NewExactMux()
	mux.Handle("/user", userMux)
	userMux.Handle("/", http.MethodGet, http.HandlerFunc(getUserHandler))
	userMux.Handle("/", http.MethodPost, http.HandlerFunc(postUserHandler))

	// Start the server using the custom ExactMux
	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

// Home page handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage!")
}

// Hello handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// Get user handler
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Retrieve user information")
}

// Post user handler
func postUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create a new user")
}
