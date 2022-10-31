package models

type Feed struct {
	Articles []Article `json:"articles"`
}

type Article struct {
	Id          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Category    string   `json:"category"`
	Rating      int      `json:"rating"`
	Authors     []string `json:"authors"`
	Content     string   `json:"content"`
}
