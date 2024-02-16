package controllers

import (
	"github.com/IlyaZh/feedsgram/internal/generated/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (p *Controller) GetSystemStatus(c *gin.Context) {
	status := api.Ok
	if p.globalState.IsFailed() {
	}
	response := api.SystemStatus{
		Status: status,
	}
	c.JSON(http.StatusOK, response)
}
