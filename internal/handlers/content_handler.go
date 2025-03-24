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

// GetHomePage godoc
// @Summary      Get Home page
// @Description  Get homepage along with child elements
// @Produce      html
// @Success      200 {string} string "HTML content"
// @Router       / [get]
func (h *ContentHandlers) GetHomePage(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "text/html")
	templates.Layout(ctx, "Home", pages.Home()).Render(ctx.Request.Context(), ctx.Writer)
}

// func (h *ContentHandlers) Layout(c *gin.Context) {
// 	c.Writer.Header().Set("Content-Type", "text/html")
// 	templates.Layout().Render(c, c.Writer) // Use the generated templ function
// }
