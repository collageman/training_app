package handlers

import (
	"encoding/json"
	"net/http"
	"user-service/models"
	"user-service/service"
)

// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user with email and password
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User Info"
// @Success 201 {object} models.User
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Error registering user"
// @Router /register [post]
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Register the user
	registeredUser, err := service.RegisterUser(user.Email, user.Password)
	if err != nil {
		http.Error(w, "Error registering user", http.StatusInternalServerError)
		return
	}

	// Return the registered user
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(registeredUser)
}

// LoginUser godoc
// @Summary Login a user
// @Description Login a user with email and password
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User Credentials"
// @Success 200 {object} models.User
// @Failure 400 {string} string "Invalid request payload"
// @Failure 401 {string} string "Invalid credentials"
// @Router /login [post]
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Authenticate the user
	authenticatedUser, err := service.LoginUser(user.Email, user.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Return the authenticated user
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authenticatedUser)
}