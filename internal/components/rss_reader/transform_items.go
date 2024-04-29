package rss_reader

import (
	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/utils"
	"github.com/mmcdole/gofeed"
)

func transformItem(sanitizer utils.Sanitizer, rawItem *gofeed.Item) entities.FeedItem {

	var image *string
	if rawItem.Image != nil {
		image = &rawItem.Image.URL
	}
	description := rawItem.Description
	if sanitizer != nil {
		description = sanitizer.SanitizeHTML(rawItem.Description)
	}
	return entities.FeedItem{
		Title:       rawItem.Title,
		Description: description,
		Content:     rawItem.Content,
		Link:        entities.Link(rawItem.Link),
		ImageURL:    image,
		PublishedAt: rawItem.PublishedParsed,
	}
}
