package message_sender

import (
	"context"
	"testing"
	"time"

	configs_mock "github.com/IlyaZh/feedsgram/internal/caches/configs/mocks"
	telegram_mock "github.com/IlyaZh/feedsgram/internal/components/telegram/mocks"
	"github.com/IlyaZh/feedsgram/internal/configs"
	"github.com/IlyaZh/feedsgram/internal/entities"
	"go.uber.org/mock/gomock"
)

func TestComponent_Start(t *testing.T) {
	ctrl := gomock.NewController(t)
	now := time.Now()

	configValue := configs.Config{
		Formatter: configs.Formatter{
			formatterFeedPost: configs.FormatterItem{
				Header: "formatted_item",
				Loop:   "",
				Footer: nil,
			},
		},
	}

	tgMock := telegram_mock.NewMockTelegram(ctrl)
	configMock := configs_mock.NewMockConfigsCache(ctrl)

	type fields struct {
		config    *configs_mock.MockConfigsCache
		telegram  *telegram_mock.MockTelegram
		input     chan []entities.FeedItem
		feedItems []entities.FeedItem
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "ok",
			fields: fields{config: configMock, telegram: tgMock, input: make(chan []entities.FeedItem), feedItems: []entities.FeedItem{
				{
					Title:       "title_1",
					Description: "description_1",
					Content:     "content_1",
					Link:        entities.Link("link_1"),
					ImageURL:    nil,
					PublishedAt: &now,
				},
			},
			},
			args: args{ctx: context.TODO()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				config:   tt.fields.config,
				telegram: tt.fields.telegram,
				input:    tt.fields.input,
			}
			c.Start(tt.args.ctx)

			tt.fields.config.EXPECT().GetValues().MaxTimes(1).Return(configValue)
			tt.fields.telegram.EXPECT().PostMessageHTML(gomock.Any(), "formatted_item").MaxTimes(1).Return(nil)

			go func() {
				tt.fields.input <- tt.fields.feedItems
				close(tt.fields.input)
			}()

			time.Sleep(time.Duration(1) * time.Second)

		})
	}
}
