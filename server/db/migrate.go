package db

import (
	"goth-todo/server/models"
)

func Migrate() {
	DB.AutoMigrate(&models.Task{})
}
