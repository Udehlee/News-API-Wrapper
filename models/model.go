package model

type Article struct {
	Author      *string `json:"author"`
	PublishedAt string  `json:"publishedAt"`
}

type NewsData struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}
