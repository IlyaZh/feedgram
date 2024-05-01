package transformer

import (
	"time"

	"github.com/IlyaZh/feedsgram/internal/entities"
)

func Feed2Source(feed entities.Feed) entities.Source {
	updatedAt := time.Now()
	return entities.Source{
		URL:         feed.FeedLink,
		Title:       &feed.Title,
		Link:        feed.Link,
		Description: feed.Description,
		IsActive:    true,
		UpdatedAt:   &updatedAt,
	}
}
