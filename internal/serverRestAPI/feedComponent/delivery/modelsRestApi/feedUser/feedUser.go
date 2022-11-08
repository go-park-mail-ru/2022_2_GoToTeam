package feedUser

type FeedUser struct {
	Articles []Article `json:"articles"`
}

type Article struct {
	Id           int       `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Tags         []string  `json:"tags"`
	Category     string    `json:"category"`
	Rating       int       `json:"rating"`
	Comments     int       `json:"comments"`
	Content      string    `json:"content"`
	CoverImgPath string    `json:"cover_img_path"`
	CoAuthor     CoAuthor  `json:"co_author"`
	Publisher    Publisher `json:"publisher"`
}

type CoAuthor struct {
	Username string `json:"username"`
	Login    string `json:"login"`
}

type Publisher struct {
	Username string `json:"username"`
	Login    string `json:"login"`
}
