package telegram

import (
	"github.com/IlyaZh/feedsgram/internal/entities"
	"mvdan.cc/xurls/v2"
)

func parseLinks(msg string) []entities.Link {
	var links []entities.Link
	parsedLinks := xurls.Strict().FindAllString(msg, -1)
	for parsedLink := range parsedLinks {
		links = append(links, entities.Link(parsedLink))
	}
	return links
}