package models

type Commentary struct {
	CommentId           int
	Content             string
	Rating              int
	Publisher           Publisher
	ArticleId           int
	CommentForCommentId string
}
