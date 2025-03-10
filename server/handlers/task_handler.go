package handlers

import (
	"goth-todo/server/models"
	"goth-todo/server/services"
	"goth-todo/server/templates"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskHandlers struct {
	Service services.TaskService
}

func NewTaskHandlers(service services.TaskService) *TaskHandlers {
	return &TaskHandlers{Service: service}
}

func (h *TaskHandlers) Foo(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/html")
	templates.Home().Render(c, c.Writer) // Use the generated templ function
}

// Loads full page with tasks
func (h *TaskHandlers) GetTasks(c *gin.Context) {
	log.Println("Getting tasks")
	tasks := h.Service.GetTasks()
	// db.DB.Find(&tasks)
	c.HTML(http.StatusOK, "tasks/task_list.html", gin.H{"tasks": tasks})
}

// Renders only the task list (HTMX)
func (h *TaskHandlers) RenderTaskList(c *gin.Context) {
	tasks := h.Service.GetTasks()
	c.HTML(http.StatusOK, "tasks/task_list.html", gin.H{"tasks": tasks})
}

// Renders the form (HTMX)
// func (h *TaskHandlers) RenderTaskForm(c *gin.Context) {
// 	c.HTML(http.StatusOK, "tasks/task_form.html", nil)
// }

// Handles adding a new task
func (h *TaskHandlers) AddTask(c *gin.Context) {
	// title := c.PostForm("title")
	// description := c.PostForm("description")
	// h.Service.AddTask(title, description)
	// h.RenderTaskList(c)
	var task models.Task
	if err := c.ShouldBind(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.Service.AddTask(&task)

	tasks := h.Service.GetTasks()
	c.HTML(http.StatusOK, "tasks/task_list.html", gin.H{"tasks": tasks})
}

// Toggles task status
func (h *TaskHandlers) ToggleTask(c *gin.Context) {
	id := c.Param("id")
	h.Service.ToggleTask(id)
	h.RenderTaskList(c)
}

// Deletes a task
func (h *TaskHandlers) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	h.Service.DeleteTask(id)
	h.RenderTaskList(c)
}
