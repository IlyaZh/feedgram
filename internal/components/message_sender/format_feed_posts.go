package message_sender

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/logger"

	"go.uber.org/zap"
)

func (c *Component) formatFeedPosts(ctx context.Context, posts []entities.FeedItem) (message entities.TelegramPost, err error) {
	log := logger.GetLoggerComponent(ctx, name)
	config := c.config.GetValues().Formatter

	formatter, ok := config[formatterFeedPost]
	if !ok {
		return "", ErrFormatterNotFound
	}

	var sb strings.Builder
	_, err = sb.WriteString(formatter.Header)
	if err != nil {
		log.Error("error while formatting feed header", zap.Error(err))
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
			log.Error("error while formatting feed item", zap.Error(err))
			return "", ErrMessageFormatFailed
		}
	}

	if formatter.Footer != nil {
		newFooter := strings.ReplaceAll(*formatter.Footer, "{{now}}", time.Now().UTC().Format(time.DateTime))
		_, err = sb.WriteString(newFooter)
		if err != nil {
			log.Error("error while formatting feed footer", zap.Error(err))
			return "", ErrMessageFormatFailed
		}
	}

	return entities.TelegramPost(sb.String()), nil
}
