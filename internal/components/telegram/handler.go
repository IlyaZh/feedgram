package telegram

import (
	"context"
	"log"

	"github.com/IlyaZh/feedsgram/internal/entities"
)

func (c *Component) handler(ctx context.Context, output chan<- entities.Message) {
	for update := range c.updates {
		if !c.messageFilter(&update) {
			log.Println("message filtered")
			continue
		}
		post := *update.Message

		links := parseLinks(post.Text)
		if len(links) == 0 {
			log.Println("links not found in message, skip")
			continue
		}

		for _, link := range links {
			output <- entities.NewMessageLink(link)
		}
	}
}
