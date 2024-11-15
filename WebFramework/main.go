package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Set up routes and handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/user", userHandler)

	// Start the server
	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
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
