// auth-service/pkg/handlers/auth_handler.go
package handlers

import (
	"auth-service/pkg/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Message string `json:"message"`
	UserID  uint   `json:"user_id"`
}

// @Summary Register new user
// @Description Register a new user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body RegisterRequest true "User registration details"
// @Success 201 {object} RegisterResponse
// @Failure 400 {object} ErrorResponse
// @Router /register [post]
func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// ErrorResponse is a struct, no need to assign it
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	user, err := services.RegisterUser(services.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	// Generate and send OTP
	otp, err := services.GenerateOTP(user.ID, "verification")
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to generate OTP"})
		return
	}

	// Send OTP via email
	go services.SendVerificationEmail(user.Email, otp.Code)

	c.JSON(http.StatusCreated, RegisterResponse{
		Message: "Registration successful. Please verify your email.",
		UserID:  user.ID,
	})
}

// @Summary Verify OTP
// @Description Verify OTP code sent to email
// @Tags auth
// @Accept json
// @Produce json
// @Param data body VerifyOTPRequest true "OTP verification data"
// @Success 200 {object} VerifyOTPResponse
// @Failure 400 {object} ErrorResponse
// @Router /verify-otp [post]
func VerifyOTP(c *gin.Context) {
	type VerifyOTPRequest struct {
		UserID string `json:"user_id"`
		Code   string `json:"code"`
	}
	var req VerifyOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// ErrorResponse is already defined, no need to redefine it
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	if err := services.VerifyOTP(req.UserID, req.Code, "verification"); err != nil {
		// No need to redefine ErrorResponse
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid or expired OTP"})
		return
	}

	type VerifyOTPResponse struct {
		Message string `json:"message"`
	}
	c.JSON(http.StatusOK, VerifyOTPResponse{
		Message: "Email verified successfully",
	})
}

// @Summary Setup MFA
// @Description Setup Multi-Factor Authentication for user
// @Tags auth
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} MFASetupResponse
// @Failure 401 {object} ErrorResponse
// @Router /setup-mfa [post]
func SetupMFA(c *gin.Context) {
	userID := c.GetUint("userID")

	secret, qrCode, err := services.SetupMFA(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	type MFASetupResponse struct {
		Secret string `json:"secret"`
		QRCode string `json:"qr_code"`
	}
	c.JSON(http.StatusOK, MFASetupResponse{
		Secret: secret,
		QRCode: qrCode,
	})
}

// @Summary Verify MFA
// @Description Verify MFA code during login
// @Tags auth
// @Accept json
// @Produce json
// @Param data body VerifyMFARequest true "MFA verification data"
// @Success 200 {object} LoginResponse
// @Failure 401 {object} ErrorResponse
// @Router /verify-mfa [post]
func VerifyMFA(c *gin.Context) {
	type VerifyMFARequest struct {
		UserID string `json:"user_id"`
		Code   string `json:"code"`
	}
	var req VerifyMFARequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	userID, err := strconv.ParseUint(req.UserID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid user ID"})
		return
	}
	tokens, err := services.VerifyMFAAndGenerateTokens(uint(userID), req.Code)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Invalid MFA code"})
		return
	}

	type LoginResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	c.JSON(http.StatusOK, LoginResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})
}

// RefreshToken handles the token refresh request
func RefreshToken(c *gin.Context) {
	// Implement the logic to refresh the token
	c.JSON(http.StatusOK, gin.H{"message": "Token refreshed successfully"})
}

func Login(c *gin.Context) {
	// Implement login logic here
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

// GetProfile handles the request to get the user's profile
func GetProfile(c *gin.Context) {
	// Implement the logic to get the user's profile
	c.JSON(http.StatusOK, gin.H{"message": "Profile retrieved successfully"})
}

// ForgotPassword handles the forgot password request
func ForgotPassword(c *gin.Context) {
	// Implement the forgot password logic here
	c.JSON(http.StatusOK, gin.H{"message": "Forgot password endpoint"})
}

// ResetPassword handles the password reset request
func ResetPassword(c *gin.Context) {
	// Implement the logic for resetting the password
	c.JSON(http.StatusOK, gin.H{"message": "Password reset successful"})
}
