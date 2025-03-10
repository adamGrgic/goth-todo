package main

import (
	"goth-todo/server/db"

	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	db.ConnectDB()
	db.MigrateDB()
}
