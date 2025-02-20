package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func LoadEnv() {
	// Try to load .env file (for local development)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	// Read environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Validate environment variables
	if dbUser == "" || dbPassword == "" || dbName == "" || dbHost == "" || dbPort == "" {
		log.Fatal("Missing required environment variables")
	}

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbUser, dbPassword, dbName, dbHost, dbPort)

	// Retry logic for database connection
	maxRetries := 5
	var err error
	for i := 0; i < maxRetries; i++ {
		DB, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Printf("Attempt %d: Failed to open database connection: %v", i+1, err)
			time.Sleep(time.Second * 5)
			continue
		}

		// Set connection pool settings
		DB.SetMaxOpenConns(25)
		DB.SetMaxIdleConns(5)
		DB.SetConnMaxLifetime(time.Minute * 5)

		// Test the connection with timeout
		err = DB.Ping()
		if err != nil {
			log.Printf("Attempt %d: Failed to ping database: %v", i+1, err)
			time.Sleep(time.Second * 5)
			continue
		}

		log.Println("Successfully connected to the database")
		return
	}

	log.Fatalf("Failed to connect to database after %d attempts: %v", maxRetries, err)
}