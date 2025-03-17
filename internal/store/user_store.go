package store

import (
	"goth-todo/internal/models"

	"github.com/gin-gonic/gin"
)

// Store user in request context
func SetUser(c *gin.Context, user *models.User) {
	c.Set("user", user)
}

// Retrieve user from request context
func GetUser(c *gin.Context) (*models.User, bool) {
	user, exists := c.Get("user")
	if !exists {
		return nil, false
	}
	return user.(*models.User), true
}
