package rss_reader

import (
	"reflect"
	"testing"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/mmcdole/gofeed"
)

func Test_transformItems(t *testing.T) {
	type args struct {
		rawItems []*gofeed.Item
	}
	tests := []struct {
		name string
		args args
		want []entities.FeedItem
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transformItems(tt.args.rawItems); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transformItems() = %v, want %v", got, tt.want)
			}
		})
	}
}
