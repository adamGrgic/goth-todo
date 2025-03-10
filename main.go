package main

import (
	"goth-todo/server/db"
	"goth-todo/server/handlers"
	"goth-todo/server/services"

	"github.com/gin-gonic/gin"
)

var statusCache map[int]string

// var db

func main() {
	db.ConnectDB()
	db.Migrate()

	taskService := services.NewGormTaskService(db.DB)
	taskHandlers := handlers.NewTaskHandlers(taskService)

	contentService := services.NewGormContentService(db.DB)
	contentHandlers := handlers.NewContentHandlers(contentService)

	r := gin.Default()
	r.LoadHTMLGlob("templ/**/*.html")

	r.Static("/static", "./static")

	r.GET("/content/home", contentHandlers.GetHomePage)
	r.GET("/task/get", taskHandlers.GetTasks)
	r.POST("/task/add", taskHandlers.AddTask)
	r.POST("/toggle/:id", taskHandlers.ToggleTask)
	r.POST("/delete/:id", taskHandlers.DeleteTask)

	r.GET("/", taskHandlers.Foo)

	r.Run(":8080")
}
