package models

type Article struct {
	ArticleId     int
	Title         string
	Description   string
	Rating        int
	CommentsCount int
	Content       string
	CoverImgPath  string

	CoAuthor     CoAuthor
	Publisher    Publisher
	CategoryName string

	Tags []string
}

type CoAuthor struct {
	Username string
	Login    string
}

type Publisher struct {
	Username string
	Login    string
}
