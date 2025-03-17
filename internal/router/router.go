package router

import (
	"goth-todo/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	taskHandlers *handlers.TaskHandler,
	contentHandlers *handlers.ContentHandlers) {

	// Serve HTML templates and static files
	// r.LoadHTMLGlob("templates/**/*.html")
	// r.Static("/public", "./public")

	// Define routes
	r.GET("/", contentHandlers.GetHomePage)

	// r.GET("/ping", systemHandlers.Ping)
	// r.GET("/content/home", contentHandlers.GetHomePage)

	// Task routes
	taskRoutes := r.Group("/task")
	{
		taskRoutes.GET("/get", taskHandlers.GetTasks)
		taskRoutes.POST("/add", taskHandlers.AddTask)
		taskRoutes.POST("/toggle/:id", taskHandlers.ToggleTask)
		taskRoutes.POST("/delete/:id", taskHandlers.DeleteTask)
	}
}
