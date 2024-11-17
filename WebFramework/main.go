package main

import (
	"fmt"
	"net/http"
	"strings"
)

type ExactMux struct {
	routes   map[string]map[string]http.Handler // path -> method -> handler
	subMuxes map[string]*ExactMux               // path -> sub-mux
}

func NewExactMux() *ExactMux {
	return &ExactMux{
		routes:   make(map[string]map[string]http.Handler),
		subMuxes: make(map[string]*ExactMux),
	}
}

// Handle registers a handler for a specific path and method.
func (e *ExactMux) Handle(path string, method string, handler http.Handler) {
	if _, exists := e.routes[path]; !exists {
		e.routes[path] = make(map[string]http.Handler)
	}
	e.routes[path][method] = handler
}

// Router registers a sub-mux for a specific path.
func (e *ExactMux) Router(path string, subMux *ExactMux) {
	e.subMuxes[path] = subMux
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

	// Check for sub-mux with exact path match
	if subMux, ok := e.subMuxes[path]; ok {
		// Adjust the request URL path
		r2 := r.Clone(r.Context())
		r2.URL.Path = "/"
		subMux.ServeHTTP(w, r2)
		return
	}

	// Check for sub-mux with path prefix match
	for subPath, subMux := range e.subMuxes {
		if strings.HasPrefix(path, subPath+"/") {
			// Adjust the request URL path
			r2 := r.Clone(r.Context())
			r2.URL.Path = strings.TrimPrefix(path, subPath)
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
	mux.Router("/user", userMux)
	userMux.Handle("/", http.MethodGet, http.HandlerFunc(getUserHandler))
	userMux.Handle("/", http.MethodPost, http.HandlerFunc(postUserHandler))
	userMux.Handle("/profile", http.MethodGet, http.HandlerFunc(userProfileHandler))

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

// User profile handler
func userProfileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User profile page")
}
