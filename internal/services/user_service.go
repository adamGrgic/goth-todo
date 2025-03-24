package services

import (
	"goth-todo/internal/models"
	"goth-todo/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUser(user *models.User, email string, password string) error
}

type UserServiceImpl struct {
	Repo repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		Repo: userRepository,
	}
}

func (r *UserServiceImpl) GetUser(user *models.User, email string, password string) error {
	// Look up user

	if err := r.Repo.GetUser(email, password); err != nil {
		return err
	}

	// Check password (assuming it's hashed)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		// c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return err
	}

	return nil
}
