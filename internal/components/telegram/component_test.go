package telegram

import (
	"context"
	"testing"
	"time"

	configsMock "github.com/IlyaZh/feedsgram/internal/caches/configs/mocks"
	metricsMock "github.com/IlyaZh/feedsgram/internal/components/metrics_storage/mocks"
	storageComponent "github.com/IlyaZh/feedsgram/internal/components/storage"
	storageMock "github.com/IlyaZh/feedsgram/internal/components/storage/mocks"
	tgAPIMock "github.com/IlyaZh/feedsgram/internal/components/telegram/mocks"
	"github.com/IlyaZh/feedsgram/internal/configs"
	"github.com/IlyaZh/feedsgram/internal/entities"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

const (
	chatID      int64 = 12
	botID       int64 = 1
	chatForFeed int64 = 42
)

func createMessage(message string, chatID int64, isCommand bool) tgbotapi.Update {
	update := tgbotapi.Update{
		ChannelPost: &tgbotapi.Message{
			Chat: &tgbotapi.Chat{
				ID: chatID,
			},
			Text: message,
		},
	}

	if isCommand {
		update.ChannelPost.Entities = make([]tgbotapi.MessageEntity, 0)
		update.ChannelPost.Entities = append(update.ChannelPost.Entities, tgbotapi.MessageEntity{
			Type:   "bot_command",
			Length: len(message),
		})
	}

	return update
}

func TestComponent_Start(t *testing.T) {
	ctrl := gomock.NewController(t)
	var limit int = 10
	var timeout int = 10

	configValue := configs.Config{
		Telegram: configs.Telegram{
			Token:      "token",
			BotID:      botID,
			UseWebhook: false,
			Limit:      &limit,
			Timeout:    &timeout,
			AllowedChatIds: map[int64]struct{}{
				chatID: {},
			},
			ChatForFeed: chatForFeed,
		},
	}

	configMock := configsMock.NewMockConfigsCache(ctrl)
	apiMock := tgAPIMock.NewMockTelegramAPI(ctrl)
	storage := storageMock.NewMockStorage(ctrl)
	metricsStorageMock := metricsMock.NewMockMetricsStorage(ctrl)

	linkAsMessage := entities.NewMessageLink("http://google.com")
	commandAsMessage := entities.NewMessageCommand("sources")

	type fields struct {
		chatID          int64
		config          *configsMock.MockConfigsCache
		api             *tgAPIMock.MockTelegramAPI
		metrics         *metricsMock.MockMetricsStorage
		updates         chan tgbotapi.Update
		storage         storageComponent.Storage
		messageText     string
		expectedMessage *entities.Message
		isCommand       bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "ok link",
			fields: fields{
				chatID:          chatID,
				config:          configMock,
				api:             apiMock,
				metrics:         metricsStorageMock,
				updates:         make(chan tgbotapi.Update, 10),
				storage:         storage,
				messageText:     "Hello http://google.com Vasya",
				expectedMessage: &linkAsMessage,
				isCommand:       false,
			},
		},
		{
			name: "ok command",
			fields: fields{
				chatID:          chatID,
				config:          configMock,
				api:             apiMock,
				metrics:         metricsStorageMock,
				updates:         make(chan tgbotapi.Update, 10),
				storage:         storage,
				messageText:     "/sources",
				expectedMessage: &commandAsMessage,
				isCommand:       true,
			},
		},
		{
			name: "not allowed chat",
			fields: fields{
				chatID:          chatID + 1,
				config:          configMock,
				api:             apiMock,
				metrics:         metricsStorageMock,
				updates:         make(chan tgbotapi.Update, 10),
				storage:         storage,
				messageText:     "Hello http://google.com Vasya",
				expectedMessage: nil,
				isCommand:       false,
			},
		},
		{
			name: "no links",
			fields: fields{
				chatID:          chatID,
				config:          configMock,
				api:             apiMock,
				metrics:         metricsStorageMock,
				updates:         make(chan tgbotapi.Update, 10),
				storage:         storage,
				messageText:     "Hello Vasya",
				expectedMessage: nil,
				isCommand:       false,
			},
		},
		{
			name: "empty message",
			fields: fields{
				chatID:          chatID,
				config:          configMock,
				api:             apiMock,
				metrics:         metricsStorageMock,
				updates:         make(chan tgbotapi.Update, 10),
				storage:         storage,
				messageText:     "",
				expectedMessage: nil,
				isCommand:       false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := make(chan entities.Message)
			defer close(ch)

			tt.fields.config.EXPECT().GetValues().AnyTimes().Return(configValue)
			tt.fields.api.EXPECT().GetUpdatesChan(gomock.Any()).Times(1).Return(tt.fields.updates)

			c := &Component{
				token:   tt.fields.config.GetValues().Telegram.Token,
				config:  tt.fields.config,
				api:     tt.fields.api,
				metrics: tt.fields.metrics,
				offset:  0}
			c.Start(context.TODO(), ch)

			tt.fields.updates <- createMessage(tt.fields.messageText, tt.fields.chatID, tt.fields.isCommand)
			if tt.fields.expectedMessage == nil {
				return
			}

			tmr := time.NewTimer(time.Duration(1) * time.Second)
			select {
			case <-tmr.C:
				t.Fail()
			case msg := <-ch:
				require.Equal(t, *tt.fields.expectedMessage, msg)
			}
		})
	}
}
