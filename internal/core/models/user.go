package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User model for authentication
type User struct {
	// ID          uint   `gorm:"primaryKey" json:"id"`
	gorm.Model
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"` // Hashed if using email/password
}

// RefreshToken model for tracking active sessions
type RefreshToken struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	Token     string    `gorm:"uniqueIndex;not null"`
	ExpiresAt time.Time `gorm:"not null"`
}
