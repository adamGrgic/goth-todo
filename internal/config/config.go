package config

import (
	"goth-todo/internal/handlers"
	"goth-todo/internal/middleware"
	"goth-todo/internal/repository"
	"goth-todo/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func RunApp(db *gorm.DB) *App {
	r := gin.Default()
	r.Static("/public", "./public")
	r.Use(middleware.Logger())
	r.Use(middleware.CORS())
	r.Use(middleware.CSPMiddleware())

	// Initialize repositories
	taskRepo := repository.NewTaskRepository(db)
	userRepo := repository.NewUserRepository(db)

	// Initialize services
	taskService := services.NewTaskService(*taskRepo)
	userService := services.NewUserService(*userRepo)

	// Initialize handlers
	taskHandlers := handlers.NewTaskHandler(taskService)
	contentHandlers := handlers.NewContentHandlers()
	userHandlers := handlers.NewUserHandler(userService)

	// might remove this section, just make this whole process one big function. Don't love the idea but we'll see.
	// Setup routes
	// router.SetupRoutes(r, taskHandlers, contentHandlers)

	r.GET("/", contentHandlers.GetHomePage)

	// User routes
	userRoutes := r.Group("/user")
	{
		userRoutes.POST("/login", userHandlers.Login)
	}

	// Task routes
	taskRoutes := r.Group("/task")
	{
		taskRoutes.GET("/get", taskHandlers.GetTasks)
		taskRoutes.POST("/add", taskHandlers.AddTask)
		taskRoutes.POST("/toggle/:id", taskHandlers.ToggleTask)
		taskRoutes.POST("/delete/:id", taskHandlers.DeleteTask)
	}
	taskRoutes.Use(middleware.JWTMiddleware())

	return &App{
		Router: r,
		DB:     db,
	}
}
