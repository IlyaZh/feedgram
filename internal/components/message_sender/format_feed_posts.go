package message_sender

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/labstack/gommon/log"
)

func (c *Component) formatFeedPosts(ctx context.Context, posts []entities.FeedItem) (message string, err error) {
	config := c.config.GetValues().Formatter

	formatter, ok := config[formatterFeedPost]
	if !ok {
		return "", ErrFormatterNotFound
	}

	var sb strings.Builder
	_, err = sb.WriteString(formatter.Header)
	if err != nil {
		log.Errorf("error while formatting feed header: %s", err.Error())
		return "", ErrMessageFormatFailed
	}

	for i, post := range posts {
		newPost := strings.ReplaceAll(formatter.Loop, "{{number}}", fmt.Sprintf("%d", i+1))
		newPost = strings.ReplaceAll(newPost, "{{link}}", string(post.Link))
		newPost = strings.ReplaceAll(newPost, "{{title}}", post.Title)
		newPost = strings.ReplaceAll(newPost, "{{description}}", post.Description)
		newPost = strings.ReplaceAll(newPost, "{{published_at}}", post.PublishedAt.UTC().Format(time.DateTime))

		_, err = sb.WriteString(newPost)
		if err != nil {
			log.Errorf("error while formatting feed item: %s", err.Error())
			return "", ErrMessageFormatFailed
		}
	}

	if formatter.Footer != nil {
		newFooter := strings.ReplaceAll(*formatter.Footer, "{{now}}", time.Now().UTC().Format(time.DateTime))
		_, err = sb.WriteString(newFooter)
		if err != nil {
			log.Errorf("error while formatting feed footer: %s", err.Error())
			return "", ErrMessageFormatFailed
		}
	}

	return sb.String(), nil
}
