package handlers

import (
	"goth-todo/internal/models"
	"goth-todo/internal/services"

	// "goth-todo/server/templates"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	TaskService services.TaskService
}

func NewTaskHandler(taskService services.TaskService) *TaskHandler {
	return &TaskHandler{
		TaskService: taskService,
	}
}

func (h *TaskHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/task/get", h.GetTasks)
	router.POST("/task/add", h.AddTask)
	router.POST("/toggle/:id", h.ToggleTask)
	router.POST("/delete/:id", h.DeleteTask)
}

// func (h *TaskHandlers) Foo(c *gin.Context) {
// 	c.Writer.Header().Set("Content-Type", "text/html")
// 	templates.Dog().Render(c, c.Writer) // Use the generated templ function
// }

// Loads full page with tasks
func (h *TaskHandler) GetTasks(c *gin.Context) {
	log.Println("Getting tasks")
	tasks, err := h.TaskService.GetTasks()
	if err != nil {
		log.Println("Something went wrong getting tasks")
	}
	// db.DB.Find(&tasks)
	// tasks_vc.Tasks().Render(c, c.Writer)
	// templates.Layout("Home", pages.Home()).Render(c, c.Writer)
	c.HTML(http.StatusOK, "tasks/task_list.html", gin.H{"tasks": tasks})

}

// Renders only the task list (HTMX)
// func (h *TaskHandler) RenderTaskList(c *gin.Context) {
// 	tasks := h.TaskService.GetTasks()
// 	c.HTML(http.StatusOK, "tasks/task_list.html", gin.H{"tasks": tasks})
// }

// Renders the form (HTMX)
// func (h *TaskHandlers) RenderTaskForm(c *gin.Context) {
// 	c.HTML(http.StatusOK, "tasks/task_form.html", nil)
// }

// Handles adding a new task
func (h *TaskHandler) AddTask(c *gin.Context) {
	// title := c.PostForm("title")
	// description := c.PostForm("description")
	// h.Service.AddTask(title, description)
	// h.RenderTaskList(c)
	var task models.Task
	if err := c.ShouldBind(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.TaskService.AddTask(&task)

	// tasks := h.TaskService.GetTasks()
	// c.HTML(http.StatusOK, "tasks/task_list.html", gin.H{"tasks": tasks})
}

// Toggles task status
func (h *TaskHandler) ToggleTask(c *gin.Context) {
	id := c.Param("id")
	h.TaskService.ToggleTask(id)
	// h.RenderTaskList(c)
}

// Deletes a task
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	// id := c.Param("id")
	// h.TaskService.DeleteTask(id)
	// h.RenderTaskList(c)
}
