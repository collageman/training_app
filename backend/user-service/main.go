// main.go
package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"user-service/handlers"
)

func main() {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	r.HandleFunc("/login", handlers.LoginUser).Methods("POST")

	// Start server
	log.Println("User Service running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}