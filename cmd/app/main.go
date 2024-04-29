package main

import (
	"context"
	"flag"
	"fmt"

	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IlyaZh/feedsgram/internal/components/message_dispatcher"
	"github.com/IlyaZh/feedsgram/internal/components/news_checker"
	"github.com/IlyaZh/feedsgram/internal/components/storage"
	"github.com/IlyaZh/feedsgram/internal/consts"
	"github.com/IlyaZh/feedsgram/internal/db"
	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/utils"
	"github.com/labstack/gommon/log"

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

	log.Info("Service initialization start")
	log.SetPrefix(consts.ServiceName)
	workEnv := os.Getenv(consts.EnvArgEnvironment)
	if workEnv == consts.EnvironmentDebug {
		log.SetLevel(log.DEBUG)
	}

	secdistPathArg := flag.String("secdist", "configs/secdist.yaml", "Set a path to secdist file relaive root dir. E.g. \"configs/secdist.yaml\"")
	configPathArg := flag.String("config", "configs/config.yaml", "Set a path to config file relaive to ./configs/ dir. E.g. \"configs/config.yaml\"")
	flag.Parse()

	configsCache := config.NewCache(ctx, *configPathArg, *secdistPathArg, time.Duration(1*time.Second))
	config := configsCache.GetValues()
	storage := storage.NewStorage(configsCache, db.CreateInstance(configsCache))
	telegram := telegram.NewTelegram(configsCache, true)

	messageBuffer := make(chan entities.Message, config.RssReader.BufferSize)
	telegram.Start(ctx, messageBuffer)

	dispatcher := message_dispatcher.NewMessageDispatcher(configsCache, storage, messageBuffer)
	dispatcher.Start(ctx)

	feedsChannel := make(chan []entities.FeedItem, config.NewsChecker.BufferSize)

	newsChecker := news_checker.NewNewsChecker(configsCache, feedsChannel, storage)
	newsCheckerPeriodc := utils.NewPeriodic("news checker", newsChecker)
	newsCheckerPeriodc.Start(ctx)

	log.Info("Service initialization has finished")

	wait()
	log.Info("Service has stopped")
}
