package service

import (
	"fmt"
	"log"

	"github.com/IlyaZh/feedsgram/controllers"
	"github.com/IlyaZh/feedsgram/internal/api"
	"github.com/IlyaZh/feedsgram/internal/models/config"
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
)

type WebServer struct {
	router  *gin.Engine
	configs *config.Cache
}

var server *WebServer

func NewWebServer(configs *config.Cache) *WebServer {
	if server != nil {
		return server
	}
	server = &WebServer{}
	server.configs = configs
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

func (s *WebServer) start() {
	settings := s.configs.GetValues().Components.Service
	server.router.Run(fmt.Sprintf("localhost:%d", settings.Port))
	log.Printf("WebServer starts at port: %d", settings.Port)
}

func (s *WebServer) Start() {
	go s.start()
}
