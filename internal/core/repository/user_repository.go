package repository

import (
	"errors"
	"fmt"
	"goth-todo/internal/core/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetUser(user *models.User, username string, password string) error {
	fmt.Println("UserRepository search user: ", username)

	err := r.DB.
		Where("email = ?", username). // assuming you're matching by email
		First(&user).Error
	fmt.Println("value of user pointer after DB find: ", user.Email)

	if err != nil {
		fmt.Println("User not found or error: ", err)
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println("Invalid password")
		return errors.New("invalid credentials")
	}

	fmt.Println("User successfully validated: ", user.Email)
	return nil
}

// func (r *UserRepository) GetUser(user models.User, email string, password string) error {
// 	// Look up user

// 	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
// 		return err
// 	}

// 	// Check password (assuming it's hashed)
// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
// 		// c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
// 		return err
// 	}
// }

// func (r *TaskRepository) GetTasks() ([]models.Task, error) {
// 	var tasks []models.Task
// 	err := r.DB.Find(&tasks).Error
// 	return tasks, err
// }

// func (r *TaskRepository) AddTask(task *models.Task) error {
// 	return r.DB.Create(task).Error
// }

// func (r *TaskRepository) ToggleTask(id string) error {
// 	var task models.Task
// 	err := r.DB.First(&task, "id = ?", id).Error
// 	if err != nil {
// 		return err
// 	}
// 	return r.DB.Save(&task).Error
// }

// func (r *TaskRepository) DeleteTask(id string) {
// 	r.DB.Delete(&models.Task{}, id)
// }
