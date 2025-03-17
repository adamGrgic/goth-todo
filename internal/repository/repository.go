package repository

import (
	"gorm.io/gorm"
)

func RegisterRepositories(db *gorm.DB) {
	NewTaskRepository(db)
}
