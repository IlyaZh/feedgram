package message_sender

import (
	"context"
	"testing"
	"time"

	configsMock "github.com/IlyaZh/feedsgram/internal/caches/configs/mocks"
	telegramMock "github.com/IlyaZh/feedsgram/internal/components/telegram/mocks"
	"github.com/IlyaZh/feedsgram/internal/configs"
	"github.com/IlyaZh/feedsgram/internal/entities"
	"go.uber.org/mock/gomock"
)

func createConfigFormatter(footer *string) configs.Config {
	return configs.Config{
		Formatter: configs.Formatter{
			formatterFeedPost: configs.FormatterItem{
				Header: "Header",
				Loop: `
{{number}}. <a href="{{link}}">{{title}}</a>
<i>Published: {{published_at}} (UTC)</i>
<b>Description:</b> {{description}}
`,
				Footer: footer,
			},
		},
	}
}

func TestComponent_formatFeedPosts(t *testing.T) {
	ctrl := gomock.NewController(t)

	footer := "Updated: {{now}}"
	config_mock := configsMock.NewMockConfigsCache(ctrl)

	tg_mock := telegramMock.NewMockTelegram(ctrl)
	testTime, _ := time.Parse(time.RFC3339, "2024-03-02T12:45:57+01:00")
	now := time.Now()
	img_1 := "image_url_1"
	posts := []entities.FeedItem{
		{
			Title:       "title_1",
			Description: "description_1",
			Content:     "content_1",
			Link:        entities.Link("link_1"),
			ImageURL:    &img_1,
			PublishedAt: &testTime,
		},
		{
			Title:       "title_2",
			Description: "description_2",
			Content:     "content_2",
			Link:        entities.Link("link_2"),
			ImageURL:    nil,
			PublishedAt: &testTime,
		},
	}

	type args struct {
		posts []entities.FeedItem
	}
	type fields struct {
		config    *configsMock.MockConfigsCache
		telegram  *telegramMock.MockTelegram
		input     <-chan []entities.FeedItem
		postsChan <-chan entities.TelegramPost
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantMessage entities.TelegramPost
		wantErr     bool
		msgFooter   *string
	}{
		{
			name:        "ok",
			fields:      fields{config: config_mock, telegram: tg_mock, input: make(chan []entities.FeedItem), postsChan: make(chan entities.TelegramPost)},
			args:        args{posts: posts},
			wantMessage: entities.TelegramPost("Header\n1. <a href=\"link_1\">title_1</a>\n<i>Published: 2024-03-02 11:45:57 (UTC)</i>\n<b>Description:</b> description_1\n\n2. <a href=\"link_2\">title_2</a>\n<i>Published: 2024-03-02 11:45:57 (UTC)</i>\n<b>Description:</b> description_2\nUpdated: " + now.UTC().Format("2006-01-02 15:04:05")),
			wantErr:     false,
			msgFooter:   &footer,
		},
		{
			name:        "ok_without_footer",
			fields:      fields{config: config_mock, telegram: tg_mock, input: make(chan []entities.FeedItem), postsChan: make(chan entities.TelegramPost)},
			args:        args{posts: posts},
			wantMessage: entities.TelegramPost("Header\n1. <a href=\"link_1\">title_1</a>\n<i>Published: 2024-03-02 11:45:57 (UTC)</i>\n<b>Description:</b> description_1\n\n2. <a href=\"link_2\">title_2</a>\n<i>Published: 2024-03-02 11:45:57 (UTC)</i>\n<b>Description:</b> description_2\n"),
			wantErr:     false,
			msgFooter:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.config.EXPECT().GetValues().MaxTimes(1).Return(createConfigFormatter(tt.msgFooter))

			c := NewMeesageSender(tt.fields.config, tt.fields.telegram, tt.fields.input, tt.fields.postsChan)
			gotMessage, err := c.formatFeedPosts(context.TODO(), tt.args.posts)
			if (err != nil) != tt.wantErr {
				t.Errorf("Component.formatFeedPosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotMessage != tt.wantMessage {
				t.Errorf("Component.formatFeedPosts() = %v, want %v", gotMessage, tt.wantMessage)
			}
		})
	}
}
