package entities

import "time"

type Feed struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Link        string     `json:"link"`
	FeedLink    string     `json:"feedLink"`
	UpdatedAt   *time.Time `json:"updatedParsed,omitempty"`
	Items       []FeedItem `json:"items"`
}

type FeedItem struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Content     string     `json:"content"`
	Link        string     `json:"link"`
	ImageURL    *string    `json:"image_url"`
	PublishedAt *time.Time `json:"publishedParsed,omitempty"`
}
