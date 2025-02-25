// auth-service/cmd/main.go
package main

import (
	"auth-service/pkg/config"
	"auth-service/pkg/handlers"

	//"auth-service/pkg/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

// @title Church Training Platform Auth API
// @version 1.0
// @description Authentication service with OTP and MFA support
// @host localhost:8080
// @BasePath /api/v1/auth
func main() {
	config.LoadConfig()
	if _, err := config.SetupDatabase(); err != nil {
		log.Fatalf("Failed to set up database: %v", err)
	}
	if _, err := config.SetupRedis(); err != nil {
		log.Fatalf("Failed to set up Redis: %v", err)
	}

	r := gin.Default()

	// Middleware
	//r.Use(middleware.CORS())

	// Auth routes
	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
		auth.POST("/verify-otp", handlers.VerifyOTP)
		//auth.POST("/setup-mfa", middleware.AuthRequired(), handlers.SetupMFA)
		auth.POST("/verify-mfa", handlers.VerifyMFA)
		auth.POST("/refresh", handlers.RefreshToken)
		auth.POST("/forgot-password", handlers.ForgotPassword)
		auth.POST("/reset-password", handlers.ResetPassword)
		//auth.GET("/profile", middleware.AuthRequired(), handlers.GetProfile)
	}

	log.Fatal(r.Run(":os.Getenv(\"PORT\")"))
}
