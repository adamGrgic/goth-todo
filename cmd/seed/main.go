package main

import (
	"errors"
	"fmt"
	"goth-todo/internal/core/models"
	db "goth-todo/internal/db"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("✅ Initializing Seeding ...")

	db.ConnectDB()

	users := []models.User{
		{Email: "alice@example.com", Password: "Blue123!"},
		{Email: "bob@example.com", Password: "Red543!$"},
	}

	for _, user := range users {
		var existing models.User
		err := db.DB.Where("email = ?", user.Email).First(&existing).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("✅ User does not exist, adding now:", user.Email)

			hashed, hashErr := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			if hashErr != nil {
				log.Fatalf("❌ Failed to hash password for %s: %v", user.Email, hashErr)
			}

			user.Password = string(hashed)

			createErr := db.DB.Create(&user).Error
			if createErr != nil {
				log.Fatalf("❌ Failed to create user %s: %v", user.Email, createErr)
			}

			log.Println("✅ Successfully added user:", user.Email)
		} else if err != nil {
			log.Fatalf("❌ Unexpected error: %v", err)
		} else {
			log.Println("⚠️  User already exists:", user.Email)
		}
	}

	fmt.Println("✅ Finished seeding users")
}
