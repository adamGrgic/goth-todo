package handlers

import (
	home_vc "goth-todo/internal/components/home"
	layout_vc "goth-todo/internal/components/layout"
	"goth-todo/internal/core/services"

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
// @Description  Example of getting a parent component and a child component
// @Produce      html
// @Success      200 {string} string "HTML content"
// @Router       / [get]
func (h *ContentHandlers) GetHomePage(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/html")

	model := layout_vc.Model{
		Context:   c,
		Title:     "Home",
		Component: home_vc.HTML(),
	}

	layout_vc.HTML(model).Render(c.Request.Context(), c.Writer)
}
