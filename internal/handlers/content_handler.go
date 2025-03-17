package handlers

import (
	"goth-todo/internal/services"
	templates "goth-todo/internal/templates/layouts"
	"goth-todo/internal/templates/pages"

	"github.com/gin-gonic/gin"
)

type ContentHandlers struct {
	Service services.ContentService
}

func NewContentHandlers() *ContentHandlers {
	return &ContentHandlers{}
}

// Loads full page with tasks
func (h *ContentHandlers) GetHomePage(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/html")

	templates.Layout("Home", pages.Home()).Render(c, c.Writer)
}

// func (h *ContentHandlers) Layout(c *gin.Context) {
// 	c.Writer.Header().Set("Content-Type", "text/html")
// 	templates.Layout().Render(c, c.Writer) // Use the generated templ function
// }
