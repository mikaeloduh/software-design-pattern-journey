package test

import (
	"fmt"
	"net/http"
)

// Handler functions remain the same
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage!")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Retrieve user information")
}

func postUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create a new user")
}

func userProfileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User profile page")
}
