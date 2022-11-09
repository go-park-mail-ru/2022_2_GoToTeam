package createArticle

type Article struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	Category      string `json:"category"`
	CoverImgPath  string `json:"cover_img_path"`
	Content       string `json:"content"`
	CoAuthorLogin string `json:"co_author_login"`

	Tags []string `json:"tags"`
}
