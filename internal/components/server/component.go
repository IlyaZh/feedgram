package server

import (
	"fmt"
	"github.com/IlyaZh/feedsgram/internal/caches/configs"
	"log"

	"github.com/IlyaZh/feedsgram/internal/controllers"
	"github.com/IlyaZh/feedsgram/internal/generated/api"
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
)

type Server struct {
	router  *gin.Engine
	configs *configs.Cache
}

var server *Server

func NewServer(configs *configs.Cache) *Server {
	if server != nil {
		return server
	}
	server = &Server{}
	server.configs = configs
	server.router = gin.Default()

	swagger, err := api.GetSwagger()
	if err != nil {
		panic(err)
	}
	server.router.Use(gin.Recovery())
	server.router.Use()

	server.router.Use(middleware.OapiRequestValidator(swagger))
	
	api.RegisterHandlersWithOptions(server.router, controllers.NewPublicApi(), api.GinServerOptions{})

	return server
}

func (s *Server) start() {
	settings := s.configs.GetValues().WebServer
	server.router.Run(fmt.Sprintf("localhost:%d", settings.Port))
	log.Printf("Server starts at port: %d", settings.Port)
}

func (s *Server) Start() {
	go s.start()
}
