package repository

import (
	"database/sql"
	"user-service/config"
	"user-service/models"
)

// CreateUser inserts a new user into the database
func CreateUser(user *models.User) error {
	query := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`
	err := config.DB.QueryRow(query, user.Email, user.Password).Scan(&user.ID)
	return err
}

// FindUserByEmail finds a user by email
func FindUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, email, password FROM users WHERE email = $1`
	err := config.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}