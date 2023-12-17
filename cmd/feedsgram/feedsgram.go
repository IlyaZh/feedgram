package main

import (
	"log"

	config "github.com/IlyaZh/feedgram/internal/components/config"
)

func main() {
	log.SetFlags(log.Llongfile | log.LUTC | log.LstdFlags)
	log.Println("Service initialization start")

	configs := config.NewCache("config.yaml", 1)
	token := configs.GetValues().Components.Telegram.Token
	log.Printf("Token from main %s", token)
	log.Println("Service initialization has finished")
	// todo
	log.Printf("Service has stopped")
}
