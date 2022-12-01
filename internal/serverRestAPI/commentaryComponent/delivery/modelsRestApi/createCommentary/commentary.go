package createCommentary

type Commentary struct {
	ArticleId           int    `json:"article_id"`
	CommentForCommentId string `json:"comment_for_comment_id"`
	Content             string `json:"content"`
}
