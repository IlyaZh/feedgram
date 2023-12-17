package main

import (
	"log"

	config "github.com/IlyaZh/feedsgram/internal/models/config"
	"github.com/IlyaZh/feedsgram/internal/models/service"
)

func main() {
	log.SetFlags(log.Llongfile | log.LUTC | log.LstdFlags)
	log.Println("Service initialization start")

	configsCache := config.NewCache("config.yaml", 1)
	config := configsCache.GetValues()
	log.Println("Service initialization has finished")
	service := service.NewWebServer(config.Components.Service.Port)
	service.Start()
	log.Printf("Service has stopped")
}
