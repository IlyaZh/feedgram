package telegram

import (
	"reflect"
	"testing"

	"github.com/IlyaZh/feedsgram/internal/caches/configs"
	"github.com/IlyaZh/feedsgram/internal/components/storage"
	"github.com/IlyaZh/feedsgram/internal/entities"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func TestComponent_Start(t *testing.T) {
	type fields struct {
		token   string
		offset  int
		config  *configs.Cache
		api     *tgbotapi.BotAPI
		isDebug bool
		updates tgbotapi.UpdatesChannel
		storage storage.Storage
		Links   chan entities.Link
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				token:   tt.fields.token,
				offset:  tt.fields.offset,
				config:  tt.fields.config,
				api:     tt.fields.api,
				isDebug: tt.fields.isDebug,
				updates: tt.fields.updates,
				storage: tt.fields.storage,
				Links:   tt.fields.Links,
			}
			c.Start()
		})
	}
}

func TestNewTelegram(t *testing.T) {
	type args struct {
		config  *configs.Cache
		isDebug bool
	}
	tests := []struct {
		name string
		args args
		want Telegram
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTelegram(tt.args.config, tt.args.isDebug); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTelegram() = %v, want %v", got, tt.want)
			}
		})
	}
}