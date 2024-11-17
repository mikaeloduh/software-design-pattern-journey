package main

import (
	"fmt"
	"net/http"
)

type ExactMux struct {
	routes map[string]http.Handler
}

// NewExactMux creates a new ExactMux instance.
func NewExactMux() *ExactMux {
	return &ExactMux{routes: make(map[string]http.Handler)}
}

// Handle registers a new route with an exact path match.
func (e *ExactMux) Handle(path string, handler http.Handler) {
	e.routes[path] = handler
}

// ServeHTTP handles incoming HTTP requests and dispatches them to the registered handlers.
func (e *ExactMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, ok := e.routes[r.URL.Path]; ok {
		handler.ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func main() {
	// Create a new ExactMux
	mux := NewExactMux()

	// Register handlers with exact path matching
	mux.Handle("/", http.HandlerFunc(homeHandler))
	mux.Handle("/hello", http.HandlerFunc(helloHandler))
	mux.Handle("/user", http.HandlerFunc(userHandler))

	// Start the server using the custom ExactMux
	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

// Home page handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Welcome to the homepage!")
}

// Hello handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// User handler
func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Retrieve user information")
	case "POST":
		fmt.Fprintf(w, "Create a new user")
	default:
		http.Error(w, "Unsupported request method", http.StatusMethodNotAllowed)
	}
}

// myHandler
func myHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w, "Welcome to the homepage!")

	case "/hello":
		fmt.Fprintf(w, "Hello, World!")

	case "/user":
		switch r.Method {
		case "GET":
			fmt.Fprintf(w, "Retrieve user information")
		case "POST":
			fmt.Fprintf(w, "Create a new user")
		default:
			http.Error(w, "Unsupported request method", http.StatusMethodNotAllowed)
		}

	default:
		http.NotFound(w, r)
	}
}
