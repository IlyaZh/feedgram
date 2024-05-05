package telegram

import (
	"context"
	"errors"
	"testing"

	configsMock "github.com/IlyaZh/feedsgram/internal/caches/configs/mocks"
	tgAPIMock "github.com/IlyaZh/feedsgram/internal/components/telegram/mocks"
	"github.com/IlyaZh/feedsgram/internal/configs"
	"github.com/IlyaZh/feedsgram/internal/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/mock/gomock"
)

func TestComponent_PostMessageHTML(t *testing.T) {
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

	type fields struct {
		config  *configsMock.MockConfigsCache
		api     *tgAPIMock.MockTelegramAPI
		updates chan tgbotapi.Update
		err     error
	}
	type args struct {
		ctx     context.Context
		message string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				config:  configMock,
				api:     apiMock,
				updates: make(chan tgbotapi.Update, 10),
				err:     nil,
			},
			args: args{
				ctx:     context.TODO(),
				message: "message to send",
			},
			wantErr: false,
		},
		{
			name: "error from API while sending",
			fields: fields{
				config:  configMock,
				api:     apiMock,
				updates: make(chan tgbotapi.Update, 10),
				err:     errors.New("error occured"),
			},
			args: args{
				ctx:     context.TODO(),
				message: "message to send",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt.fields.config.EXPECT().GetValues().AnyTimes().Return(configValue)
		tt.fields.api.EXPECT().Send(utils.CreateTelegramHTMLMessage(chatForFeed, tt.args.message)).Times(1).Return(tgbotapi.Message{}, tt.fields.err)

		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				token:  tt.fields.config.GetValues().Telegram.Token,
				config: tt.fields.config,
				api:    tt.fields.api,
				offset: 0}
			if err := c.PostMessageHTML(tt.args.ctx, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("Component.PostMessageHTML() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
