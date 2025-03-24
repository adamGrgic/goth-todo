package config

import (
	"goth-todo/internal/handlers"
	"goth-todo/internal/middleware"
	"goth-todo/internal/repository"
	"goth-todo/internal/router"
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

	// Initialize services
	taskService := services.NewTaskService(*taskRepo)

	// Initialize handlers
	taskHandlers := handlers.NewTaskHandler(taskService)
	contentHandlers := handlers.NewContentHandlers()

	// Setup routes
	router.SetupRoutes(r, taskHandlers, contentHandlers)

	return &App{
		Router: r,
		DB:     db,
	}
}
