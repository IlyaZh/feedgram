package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IlyaZh/feedsgram/internal/components/message_dispatcher"
	"github.com/IlyaZh/feedsgram/internal/components/storage"
	"github.com/IlyaZh/feedsgram/internal/db"
	"github.com/IlyaZh/feedsgram/internal/entities"

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
	ctx := context.TODO()
	log.SetFlags(log.Llongfile | log.LUTC | log.LstdFlags | log.Ldate | log.Ltime)
	log.Println("Service initialization start")

	configsCache := config.NewCache(ctx, "config.yaml", time.Duration(1*time.Second))
	log.Println("Service initialization has finished")
	storage := storage.NewStorage(configsCache, db.CreateInstance(configsCache))
	telegram := telegram.NewTelegram(configsCache, true)

	messageBuffer := make(chan entities.Message, configsCache.GetValues().RssReader.BufferSize)
	telegram.Start(ctx, messageBuffer)

	dispatcher := message_dispatcher.NewMessageDispatcher(configsCache, storage, messageBuffer)
	dispatcher.Start(ctx)

	wait()
	log.Println("Service has stopped")
}
