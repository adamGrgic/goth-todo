package db

import (
	"goth-todo/internal/models"
)

func Migrate() {
	DB.AutoMigrate(&models.Task{})
}
