package db

import (
	"goth-todo/internal/core/models"
)

func Migrate() {
	DB.AutoMigrate(&models.Task{})
}
