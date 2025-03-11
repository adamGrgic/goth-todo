package config

import (
	"goth-todo/server/handlers"
	"goth-todo/server/middleware"
	"goth-todo/server/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	Router         *gin.Engine
	TaskHandler    *handlers.TaskHandlers
	ContentHandler *handlers.ContentHandlers
	SystemHandler  *handlers.SystemHandlers
}

func CreateApp(db *gorm.DB) *App {
	app := setupServicesAndHandlers(db)

	app.setupRoutes()
	return app
}

func setupServicesAndHandlers(db *gorm.DB) *App {
	taskService := services.NewGormTaskService(db)
	taskHandlers := handlers.NewTaskHandlers(taskService)

	contentService := services.NewGormContentService(db)
	contentHandlers := handlers.NewContentHandlers(contentService)

	systemService := services.NewGormSystemService(db)
	systemHandlers := handlers.NewSystemHandlers(systemService)

	app := &App{
		Router:         gin.Default(),
		TaskHandler:    taskHandlers,
		ContentHandler: contentHandlers,
		SystemHandler:  systemHandlers,
	}

	return app
}

func (a *App) setupRoutes() {

	a.Router.LoadHTMLGlob("templ/**/*.html")
	a.Router.Static("/public", "./public")

	a.Router.GET("/", a.SystemHandler.Ping)
	a.Router.GET("/content/home", a.ContentHandler.GetHomePage)
	a.Router.GET("/task/get", a.TaskHandler.GetTasks)
	a.Router.POST("/task/add", a.TaskHandler.AddTask)
	a.Router.POST("/toggle/:id", a.TaskHandler.ToggleTask)
	a.Router.POST("/delete/:id", a.TaskHandler.DeleteTask)
	a.Router.GET("/home", a.ContentHandler.Foo)
}

func (a *App) setupMiddleware() {
	a.Router.Use(middleware.Logger())
}
