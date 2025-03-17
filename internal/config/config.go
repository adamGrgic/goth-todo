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

// type App struct {
// 	Router         *gin.Engine
// 	TaskHandler    *handlers.TaskHandlers
// 	ContentHandler *handlers.ContentHandlers
// 	SystemHandler  *handlers.SystemHandlers
// }

//	type App struct {
//		Router         *gin.Engine
//		TaskHandler    *handlers.TaskHandlers
//		ContentHandler *handlers.ContentHandlers
//		SystemHandler  *handlers.SystemHandlers
//	}
type App struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func CreateApp(db *gorm.DB) *App {
	r := gin.Default()
	r.Use(middleware.Logger())

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

// func setupServicesAndHandlers(db *gorm.DB) *App {
// 	taskService := services.NewTaskService()
// 	contentService := services.NewGormContentService(db)
// 	systemService := services.NewGormSystemService(db)

// 	taskHandlers := handlers.NewTaskHandlers(taskService)
// 	contentHandlers := handlers.NewContentHandlers(contentService)

// 	systemHandlers := handlers.NewSystemHandlers(systemService)

// 	app := &App{
// 		Router:         gin.Default(),
// 		TaskHandler:    taskHandlers,
// 		ContentHandler: contentHandlers,
// 		SystemHandler:  systemHandlers,
// 	}

// 	return app
// }

// func (a *App) setupRoutes() {

// 	a.Router.LoadHTMLGlob("templ/**/*.html")
// 	a.Router.Static("/public", "./public")

// 	a.Router.GET("/", a.ContentHandler.Layout)
// 	a.Router.GET("/ping", a.SystemHandler.Ping)
// 	a.Router.GET("/content/home", a.ContentHandler.GetHomePage)
// 	a.Router.GET("/task/get", a.TaskHandler.GetTasks)
// 	a.Router.POST("/task/add", a.TaskHandler.AddTask)
// 	a.Router.POST("/toggle/:id", a.TaskHandler.ToggleTask)
// 	a.Router.POST("/delete/:id", a.TaskHandler.DeleteTask)
// }

// func (a *App) setupMiddleware() {
// 	a.Router.Use(middleware.Logger())
// }

// type App struct {
// 	Router   *gin.Engine
// 	Handlers map[string]handlers.Handler
// }

// func CreateApp(db *gorm.DB) *App {
// 	app := &App{
// 		Router:   gin.Default(),
// 		Handlers: make(map[string]handlers.Handler),
// 	}

// 	app.setupServicesAndHandlers(db)
// 	app.setupMiddleware()
// 	app.setupRoutes()

// 	return app
// }

// func (a *App) setupServicesAndHandlers(db *gorm.DB) {
// 	// Automatically register services and handlers
// 	serviceRegistry := map[string]services.Service{
// 		"content": services.NewGormContentService(db),
// 	}

// 	for key, service := range serviceRegistry {
// 		a.Handlers[key] = handlers.NewHandler(service)
// 	}
// }

// func (a *App) setupRoutes() {
// 	for _, handler := range a.Handlers {
// 		handler.RegisterRoutes(a.Router)
// 	}
// }

// func (a *App) setupMiddleware() {
// 	a.Router.Use(middleware.Logger())
// }

// func setupServicesAndHandlers(db *gorm.DB) *App {
// 	taskService := services.NewGormTaskService(db)
// 	taskHandlers := handlers.NewTaskHandlers(taskService)

// 	contentService := services.NewGormContentService(db)
// 	contentHandlers := handlers.NewContentHandlers(contentService)

// 	systemService := services.NewGormSystemService(db)
// 	systemHandlers := handlers.NewSystemHandlers(systemService)

// 	app := &App{
// 		Router:         gin.Default(),
// 		TaskHandler:    taskHandlers,
// 		ContentHandler: contentHandlers,
// 		SystemHandler:  systemHandlers,
// 	}

// 	return app
// }

// func (a *App) setupRoutes() {

// 	a.Router.LoadHTMLGlob("templ/**/*.html")
// 	a.Router.Static("/public", "./public")

// 	a.Router.GET("/", a.ContentHandler.Layout)
// 	a.Router.GET("/ping", a.SystemHandler.Ping)
// 	a.Router.GET("/content/home", a.ContentHandler.GetHomePage)
// 	a.Router.GET("/task/get", a.TaskHandler.GetTasks)
// 	a.Router.POST("/task/add", a.TaskHandler.AddTask)
// 	a.Router.POST("/toggle/:id", a.TaskHandler.ToggleTask)
// 	a.Router.POST("/delete/:id", a.TaskHandler.DeleteTask)
// }

// func (a *App) setupMiddleware() {
// 	a.Router.Use(middleware.Logger())
// }
