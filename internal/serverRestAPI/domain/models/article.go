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
	Email    string
	Username string
	Login    string
}

type Publisher struct {
	Email    string
	Username string
	Login    string
}
