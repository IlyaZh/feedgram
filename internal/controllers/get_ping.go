package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (p *Controller) GetPing(c *gin.Context) {
	status := http.StatusOK
	if p.globalState.IsFailed() {
		status = http.StatusInternalServerError
	}
	c.Status(status)

}
