package handlers

import (
	"goth-todo/server/services"
	"net/http"

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
