package main

import (
	"goth-todo/server/config"
	"goth-todo/server/db"
)

var statusCache map[int]string

// func configureApp(db *gorm.DB) *App {
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

// 	app.setupRoutes()
// 	return app
// }

// func (a *App) setupRoutes() {

// 	a.Router.LoadHTMLGlob("templ/**/*.html")
// 	a.Router.Static("/static", "./static")

// 	a.Router.GET("/", middleware.Logger(), a.SystemHandler.Ping)
// 	a.Router.GET("/content/home", a.ContentHandler.GetHomePage)
// 	a.Router.GET("/task/get", a.TaskHandler.GetTasks)
// 	a.Router.POST("/task/add", a.TaskHandler.AddTask)
// 	a.Router.POST("/toggle/:id", a.TaskHandler.ToggleTask)
// 	a.Router.POST("/delete/:id", a.TaskHandler.DeleteTask)
// }

// sample middleware
// func Logger() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		t := time.Now()

// 		// Set example variable
// 		c.Set("example", "12345")

// 		// before request

// 		c.Next()

// 		// after request
// 		latency := time.Since(t)
// 		log.Print(latency)

// 		// access the status we are sending
// 		status := c.Writer.Status()
// 		log.Println(status)
// 	}
// }

func main() {
	db.ConnectDB()
	db.Migrate()

	app := config.CreateApp(db.DB)
	// app.Router.Use(Logger())
	app.Router.Run(":9090")
}

// func main() {
// 	db.ConnectDB()
// 	db.Migrate()

// 	taskService := services.NewGormTaskService(db.DB)
// 	taskHandlers := handlers.NewTaskHandlers(taskService)

// 	contentService := services.NewGormContentService(db.DB)
// 	contentHandlers := handlers.NewContentHandlers(contentService)

// 	r := gin.Default()
// 	r.LoadHTMLGlob("templ/**/*.html")

// 	r.Static("/static", "./static")

// 	r.GET("/content/home", contentHandlers.GetHomePage)
// 	r.GET("/task/get", taskHandlers.GetTasks)
// 	r.POST("/task/add", taskHandlers.AddTask)
// 	r.POST("/toggle/:id", taskHandlers.ToggleTask)
// 	r.POST("/delete/:id", taskHandlers.DeleteTask)

// 	r.GET("/", taskHandlers.Foo)

// 	r.Run(":8080")
// }
