package service

import (
	"fmt"

	"github.com/IlyaZh/feedsgram/controllers"
	"github.com/IlyaZh/feedsgram/internal/api"
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
)

type WebServer struct {
	port   int
	router *gin.Engine
}

var server *WebServer

func NewWebServer(port int) *WebServer {
	if server != nil {
		return server
	}
	server = &WebServer{}
	server.port = port
	server.router = gin.Default()

	swagger, err := api.GetSwagger()
	if err != nil {
		panic(err)
	}
	server.router.Use(gin.Recovery())
	server.router.Use()

	server.router.Use(middleware.OapiRequestValidator(swagger))

	publicApi := controllers.NewPublicApi()
	api.RegisterHandlersWithOptions(server.router, publicApi, api.GinServerOptions{})

	return server
}

func (s *WebServer) Start() {
	server.router.Run(fmt.Sprintf("localhost:%d", server.port))
}
