package gonewsapi

type Response struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

type Article struct {
	Source      []Sourcing `json:"source"`
	Author      string     `json:"author"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Url         string     `json:"url"`
	PublishedAt string     `json:"publishedAt"`
	Content     string     `json:"content"`
}

type Sourcing struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
