package handlers

import (
	"goth-todo/server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SystemHandlers struct {
	Service services.SystemService
}

func NewSystemHandlers(service services.SystemService) *SystemHandlers {
	return &SystemHandlers{Service: service}
}

func (h *SystemHandlers) Ping(c *gin.Context) {
	response := h.Service.Ping()
	c.JSON(http.StatusOK, gin.H{
		"message": response,
	})
}
