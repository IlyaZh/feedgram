package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	config "github.com/IlyaZh/feedsgram/internal/models/config"
	"github.com/IlyaZh/feedsgram/internal/models/service"
	"github.com/IlyaZh/feedsgram/internal/models/telegram"
)

func wait() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Printf("Signal recevied %d", sig)
		fmt.Println()
		done <- true
	}()
	<-done
}

func main() {
	log.SetFlags(log.Llongfile | log.LUTC | log.LstdFlags)
	log.Println("Service initialization start")

	configsCache := config.NewCache("config.yaml", 1)
	config := configsCache.GetValues()
	log.Println("Service initialization has finished")
	// TODO change configs values to config cache component
	log.Println("Telegram start")
	telegram := telegram.NewTelegram(config.Components.Telegram, true)
	go telegram.Start()
	log.Println("Telegram has started")
	log.Println("WebServer start")
	service := service.NewWebServer(config.Components.Service.Port)
	go service.Start()
	log.Println("WebServer has started")

	wait()
	log.Println("Service has stopped")
}
