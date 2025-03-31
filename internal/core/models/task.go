package models

import "gorm.io/gorm"

type Task struct {
	// ID          uint   `gorm:"primaryKey" json:"id"`
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
