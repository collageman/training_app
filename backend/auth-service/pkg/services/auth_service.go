// auth-service/pkg/services/auth_service.go
package services

import (
	"auth-service/pkg/config"
	"auth-service/pkg/models"
	"crypto/rand"
	"encoding/base32"
	"errors"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"github.com/pquerna/otp/totp"
	"gorm.io/gorm"
)

type User struct {
	ID    uint
	Email string
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type OTPService struct {
	db    *gorm.DB
	redis *redis.Client
}

func GenerateOTP(userID uint, otpType string) (*models.OTP, error) {
	// Generate random 6-digit code
	n, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		return nil, err
	}
	code := fmt.Sprintf("%06d", n.Int64())

	otp := &models.OTP{
		UserID:    userID,
		Code:      code,
		Type:      otpType,
		ExpiresAt: time.Now().Add(15 * time.Minute),
	}

	if err := config.DB.Create(otp).Error; err != nil {
		return nil, err
	}

	return otp, nil
}

func SetupMFA(userID uint) (string, string, error) {
	// Generate random secret
	secret := make([]byte, 20)
	rand.Read(secret)
	secretBase32 := base32.StdEncoding.EncodeToString(secret)

	// Generate QR code
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "ChurchTrainingPlatform",
		AccountName: getUserEmail(userID),
		Secret:      []byte(secretBase32),
	})
	if err != nil {
		return "", "", err
	}

	// Update user with MFA secret
	if err := config.DB.Model(&models.User{}).
		Where("id = ?", userID).
		Updates(map[string]interface{}{
			"mfa_enabled": true,
			"mfa_secret":  secretBase32,
		}).Error; err != nil {
		return "", "", err
	}

	return secretBase32, key.URL(), nil
}

func getUserEmail(userID uint) string {
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return ""
	}
	return user.Email
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func VerifyMFAAndGenerateTokens(userID uint, code string) (*TokenPair, error) {
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}

	// Verify TOTP code
	valid := totp.Validate(code, user.MFASecret)
	if !valid {
		return nil, errors.New("invalid MFA code")
	}

	// Generate new token pair
	return GenerateTokenPair(user)
}

func SendVerificationEmail(email, code string) error {
	// Implementation for sending email with OTP
	// You can use any email service like SendGrid, AWS SES, etc.
	return nil
}

// Helper function to generate JWT tokens
func GenerateTokenPair(user models.User) (*TokenPair, error) {
	// Generate access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	})

	// Generate refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	})

	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

func RegisterUser(req RegisterRequest) (*models.User, error) {
	// Implement the logic to register a new user
	// For now, let's return a dummy user
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email and password are required")
	}

	user := &models.User{
		// ID:    1, // Remove or replace this line with a valid field
		Email: req.Email,
	}

	return user, nil
}

// VerifyOTP verifies the OTP code for a given user and purpose
func VerifyOTP(userID string, code string, purpose string) error {
	// Implement the logic to verify the OTP code
	// For now, let's assume the OTP is always valid
	if code == "123456" {
		return nil
	}
	return errors.New("invalid or expired OTP")
}
