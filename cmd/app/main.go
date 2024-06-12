package main

import (
	"context"
	"flag"

	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IlyaZh/feedsgram/internal/components/message_dispatcher"
	"github.com/IlyaZh/feedsgram/internal/components/message_sender"
	"github.com/IlyaZh/feedsgram/internal/components/news_checker"
	"github.com/IlyaZh/feedsgram/internal/components/storage"
	"github.com/IlyaZh/feedsgram/internal/components/telegram"
	"github.com/IlyaZh/feedsgram/internal/consts"
	"github.com/IlyaZh/feedsgram/internal/db"
	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/logger"
	"github.com/IlyaZh/feedsgram/internal/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jamillosantos/logctx"
	"go.uber.org/zap"

	config "github.com/IlyaZh/feedsgram/internal/caches/configs"
)

func wait(ctx context.Context) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		logctx.Info(ctx, "OS Signal recevied", zap.String("signal", sig.String()))
		done <- true
	}()
	<-done
}

func main() {
	workEnv := os.Getenv(consts.EnvArgEnvironment)
	isDebug := (workEnv == consts.EnvironmentDebug)

	ctx := context.TODO()
	ctx = logctx.WithLogger(ctx, logger.NewLogger(ctx, isDebug).With(zap.String("service", consts.ServiceName)))

	logctx.Info(ctx, "Service initialization start")

	defer func(ctx context.Context) {
		if r := recover(); r != nil {
			logctx.Error(ctx, "panic occured", zap.Any("panic", r))
			panic(r)
		}
	}(ctx)

	secdistPathArg := flag.String("secdist", "configs/secdist.yaml", "Set a path to secdist file relaive root dir. E.g. \"configs/secdist.yaml\"")
	configPathArg := flag.String("config", "configs/config.yaml", "Set a path to config file relaive to root dir. E.g. \"configs/config.yaml\"")
	flag.Parse()

	configsCache := config.NewCache(ctx, *configPathArg, *secdistPathArg, time.Duration(5*time.Second))
	config := configsCache.GetValues()
	storage := storage.NewStorage(configsCache, db.CreateInstance(ctx, configsCache))

	tgBot, err := tgbotapi.NewBotAPI(config.Telegram.Token)
	if err != nil {
		panic(err)
	}
	telegram := telegram.NewTelegram(configsCache, tgBot)

	messageBuffer := make(chan entities.Message, config.RssReader.BufferSize)
	telegram.Start(ctx, messageBuffer)

	postsChannel := make(chan entities.TelegramPost)

	dispatcher := message_dispatcher.NewMessageDispatcher(configsCache, storage, messageBuffer)
	dispatcher.Start(ctx)

	feedsChannel := make(chan []entities.FeedItem, config.NewsChecker.BufferSize)

	newsCheckerPeriodc := utils.NewPeriodic("news checker", news_checker.NewNewsChecker(configsCache, feedsChannel, storage))
	newsCheckerPeriodc.Start(ctx)

	sender := message_sender.NewMeesageSender(configsCache, telegram, feedsChannel, postsChannel)
	sender.Start(ctx)

	logctx.Info(ctx, "Service initialization has finished")

	if config.Telegram.MessageWhenStart {
		_ = telegram.PostMessageHTML(ctx, entities.TelegramPost("Feedgram has started"))
	}

	wait(ctx)
	logctx.Info(ctx, "Service has stopped")
}
