package models

import (
    "time"
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Email         string `gorm:"uniqueIndex;not null"`
    Password      string `gorm:"not null"`
    FirstName     string
    LastName      string
    Role          string `gorm:"default:'user'"`
    MFAEnabled    bool   `gorm:"default:false"`
    MFASecret     string
    LastLogin     time.Time
    EmailVerified bool   `gorm:"default:false"`
}

type OTP struct {
    gorm.Model
    UserID    uint
    Code      string
    Type      string // "verification", "password-reset", "mfa"
    ExpiresAt time.Time
}

type Session struct {
    gorm.Model
    UserID       uint
    Token        string `gorm:"uniqueIndex"`
    RefreshToken string `gorm:"uniqueIndex"`
    ExpiresAt    time.Time
    Device       string
    IP           string
}