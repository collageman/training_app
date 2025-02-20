package service

import (
	"user-service/models"
	"user-service/repository"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser handles user registration
func RegisterUser(email, password string) (*models.User, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create the user
	user := &models.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	// Save the user to the database
	err = repository.CreateUser(user)
	return user, err
}

// LoginUser handles user login
func LoginUser(email, password string) (*models.User, error) {
	// Find the user by email
	user, err := repository.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	// Compare passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}