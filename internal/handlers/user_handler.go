package handlers

import (
	"goth-todo/internal/auth"
	"goth-todo/internal/models"
	"goth-todo/internal/services"
	templates "goth-todo/internal/templates/layouts"
	"goth-todo/internal/templates/pages"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

// func (h *UserHandler) RegisterRoutes(router *gin.Engine) {
// 	// router.GET("/task/get", h.GetTasks)
// 	// router.GET("/user/login/", h.Login)
// 	router.POST("/user/login/", h.Login)
// 	// router.POST("/delete/:id", h.DeleteTask)
// }

// func (h *UserHandler) LoginPage(c *gin.Context) {

// }

func (h *UserHandler) Login(c *gin.Context) {
	// Parse form values
	email := c.PostForm("email")
	password := c.PostForm("password")

	var user models.User

	h.UserService.GetUser(&user, email, password)

	if user.Email == "" {

	}

	// Generate JWT
	token, err := auth.GenerateJWT(user.Email)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.Set("jwt_token", token)

	templates.Layout(c, "Home", pages.Home()).Render(c.Request.Context(), c.Writer)
}

// Return token (as JSON or set as cookie)
// c.JSON(http.StatusOK, gin.H{
// 	"token": token,
// 	"user": gin.H{
// 		"id":    user.ID,
// 		"email": user.Email,
// 	},
// })
