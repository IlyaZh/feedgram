package rss_reader

import (
	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/mmcdole/gofeed"
)

func transformItems(rawItems []*gofeed.Item) []entities.FeedItem {
	items := make([]entities.FeedItem, 0, len(rawItems))

	for _, rawItem := range rawItems {
		var image *string
		if rawItem.Image != nil {
			image = &rawItem.Image.URL
		}
		item := entities.FeedItem{
			Title:       rawItem.Title,
			Description: rawItem.Description,
			Content:     rawItem.Content,
			Link:        rawItem.Link,
			ImageURL:    image,
			PublishedAt: rawItem.PublishedParsed,
		}
		items = append(items, item)
	}

	return items
}
