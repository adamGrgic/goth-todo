package handlers

import (
	"goth-todo/internal/core/services"
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

// GetHomePage godoc
// @Summary      Get Home page
// @Description  Get homepage along with child elements
// @Produce      html
// @Success      200 {string} string "HTML content"
// @Router       / [get]
func (h *ContentHandlers) GetHomePage(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/html")
	templates.Layout(c, "Home", pages.Home()).Render(c.Request.Context(), c.Writer)
}

// func (h *ContentHandlers) Layout(c *gin.Context) {
// 	c.Writer.Header().Set("Content-Type", "text/html")
// 	templates.Layout().Render(c, c.Writer) // Use the generated templ function
// }
