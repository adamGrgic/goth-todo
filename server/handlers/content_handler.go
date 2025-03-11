package handlers

import (
	"goth-todo/server/services"
	"net/http"

	"goth-todo/server/templates"

	"github.com/gin-gonic/gin"
)

type ContentHandlers struct {
	Service services.ContentService
}

func NewContentHandlers(service services.ContentService) *ContentHandlers {
	return &ContentHandlers{Service: service}
}

// Loads full page with tasks
func (h *ContentHandlers) GetHomePage(c *gin.Context) {
	// tasks := h.Service.GetTasks()
	// db.DB.Find(&tasks)
	c.HTML(http.StatusOK, "layouts/base.html", gin.H{})
}

func (h *ContentHandlers) Foo(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/html")
	templates.Layout().Render(c, c.Writer) // Use the generated templ function
}
