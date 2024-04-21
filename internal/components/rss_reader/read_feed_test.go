package rss_reader

import (
	"context"
	"reflect"
	"testing"

	"github.com/IlyaZh/feedsgram/internal/entities"
)

func TestComponent_ReadFeed(t *testing.T) {
	type args struct {
		ctx  context.Context
		link entities.Link
	}
	tests := []struct {
		name    string
		c       *Component
		args    args
		want    entities.Feed
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.ReadFeed(tt.args.ctx, tt.args.link)
			if (err != nil) != tt.wantErr {
				t.Errorf("Component.ReadFeed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Component.ReadFeed() = %v, want %v", got, tt.want)
			}
		})
	}
}
