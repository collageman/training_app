package main

import (
	"log"
	"net/http"
	"user-service/config"
	"user-service/handlers"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize the router
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	r.HandleFunc("/login", handlers.LoginUser).Methods("POST")

	// Swagger documentation
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Start server
	log.Println("User Service running on port 8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}
