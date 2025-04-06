package handlers

import (
	"fmt"
	"goth-todo/internal/auth"
	login_vc "goth-todo/internal/components/login"
	"goth-todo/internal/core/models"
	"goth-todo/internal/core/services"
	"net/http"

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
	fmt.Println("Login method running")

	email := c.PostForm("email")
	password := c.PostForm("password")

	fmt.Println("form email: ", email)
	fmt.Println("form password: ", password)

	var user models.User
	h.UserService.GetUser(&user, email, password)
	fmt.Println("retrieved user: ", user.Email)

	if user.Email == "" {
		fmt.Println("User not found, please try again.")
		model := login_vc.Model{
			Username: email,
			ErrorMsg: "User not found, please try again.",
		}
		login_vc.HTML(model).Render(c.Request.Context(), c.Writer)
		return
	}

	// ✅ Generate JWT
	err := auth.SetUserJWT(c, user.Email)
	if err != nil {
		c.String(http.StatusInternalServerError, "Could not generate token")
		return
	}

	// ✅ Redirect with HTMX
	c.Header("HX-Redirect", "/")
	c.Status(http.StatusOK)
}

func (h *UserHandler) Register(c *gin.Context) {
	fmt.Println("Register method running")

}

// Return token (as JSON or set as cookie)
// c.JSON(http.StatusOK, gin.H{
// 	"token": token,
// 	"user": gin.H{
// 		"id":    user.ID,
// 		"email": user.Email,
// 	},
// })
