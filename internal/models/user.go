package models

import (
	"time"

	"github.com/google/uuid"
)

// User model for authentication
type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email    string    `gorm:"uniqueIndex;not null"`
	Password string    `gorm:"not null"` // Hashed if using email/password
}

// RefreshToken model for tracking active sessions
type RefreshToken struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	Token     string    `gorm:"uniqueIndex;not null"`
	ExpiresAt time.Time `gorm:"not null"`
}
