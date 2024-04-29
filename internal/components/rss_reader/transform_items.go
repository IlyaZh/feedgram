package rss_reader

import (
	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/mmcdole/gofeed"
)

func transformItem(rawItem *gofeed.Item) entities.FeedItem {

	var image *string
	if rawItem.Image != nil {
		image = &rawItem.Image.URL
	}
	return entities.FeedItem{
		Title:       rawItem.Title,
		Description: rawItem.Description,
		Content:     rawItem.Content,
		Link:        entities.Link(rawItem.Link),
		ImageURL:    image,
		PublishedAt: rawItem.PublishedParsed,
	}
}
