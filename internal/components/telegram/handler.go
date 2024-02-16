package telegram

import (
	"context"
	"log"
)

func (c *Component) handler(ctx context.Context) {
	for update := range c.updates {
		if !c.messageFilter(&update) {
			log.Println("message filtered")
			continue
		}
		post := update.ChannelPost

		links := parseLinks(post.Text)
		if len(links) == 0 {
			log.Println("links not found in message, skip")
			continue
		}
		for _, link := range links {
			_, err := c.rssReader.ReadFeed(ctx, link)
			if err != nil {
				log.Printf("Eror while read feed: %s. Error: %s\n", link, err.Error())
				continue
			}
		}

	}
}
