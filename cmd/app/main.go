package main

import (
	"fmt"
	"github.com/IlyaZh/feedsgram/internal/components/global_state"
	"github.com/IlyaZh/feedsgram/internal/components/rss_reader"
	"github.com/IlyaZh/feedsgram/internal/components/server"
	"github.com/IlyaZh/feedsgram/internal/components/storage"
	"github.com/IlyaZh/feedsgram/internal/db"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/IlyaZh/feedsgram/internal/caches/configs"
	"github.com/IlyaZh/feedsgram/internal/components/telegram"
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
	log.SetFlags(log.Llongfile | log.LUTC | log.LstdFlags | log.Ldate | log.Ltime)
	log.Println("Service initialization start")

	globalState := globalstate.NewGlobalState()
	configsCache := config.NewCache("config.yaml", time.Duration(1*time.Second))
	log.Println("Service initialization has finished")
	storage := storage.NewStorage(configsCache, db.CreateInstance(configsCache))
	rssReader := rss_reader.NewRssReader(configsCache)
	log.Println("Telegram component start")
	telegram := telegram.NewTelegram(configsCache, storage, rssReader, true)
	telegram.Start()
	log.Println("Telegram component has started")
	log.Println("Server start")
	server := server.NewServer(configsCache)
	server.Start()
	log.Println("Server has started")

	globalState.SetFailed(false)

	wait()
	log.Println("Service has stopped")
}
