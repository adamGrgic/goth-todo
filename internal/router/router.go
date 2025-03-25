package router

// import (
// 	"goth-todo/internal/handlers"
// 	"goth-todo/internal/middleware"

// 	"github.com/gin-gonic/gin"
// )

// possibly remove

// TODO: make slice of generics, if possible
// func SetupRoutes(
// 	r *gin.Engine,
// 	taskHandlers *handlers.TaskHandler,
// 	contentHandlers *handlers.ContentHandlers,
// 	userHandlers) {

// 	// Define routes
// 	r.GET("/", contentHandlers.GetHomePage)

// 	// Task routes
// 	taskRoutes := r.Group("/task")
// 	{
// 		taskRoutes.GET("/get", taskHandlers.GetTasks)
// 		taskRoutes.POST("/add", taskHandlers.AddTask)
// 		taskRoutes.POST("/toggle/:id", taskHandlers.ToggleTask)
// 		taskRoutes.POST("/delete/:id", taskHandlers.DeleteTask)
// 	}
// 	taskRoutes.Use(middleware.JWTMiddleware())
// }
