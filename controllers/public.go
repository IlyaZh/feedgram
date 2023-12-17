package controllers

import (
	"net/http"

	"github.com/IlyaZh/feedsgram/internal/api"
	"github.com/gin-gonic/gin"
)

type Public struct{}

var public *Public

func NewPublicApi() *Public {
	if public != nil {
		return public
	}
	public = &Public{}
	return public
}

func (p *Public) GetPing(c *gin.Context) {
	c.Status(http.StatusOK)
	// c.Status(http.StatusInternalServerError)
}

func (p *Public) GetSystemStatus(c *gin.Context) {
	response := api.SystemStatus{}
	c.JSON(http.StatusOK, response)
}
