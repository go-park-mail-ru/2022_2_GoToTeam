package updateArticle

type Article struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Content     string `json:"content"`

	Tags []string `json:"tags"`
}
