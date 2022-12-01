package getAllCommentariesForArticle

type CommentariesForArticle struct {
	Commentaries []Commentary `json:"commentaries"`
}

type Commentary struct {
	CommentId           int       `json:"comment_id"`
	Content             string    `json:"content"`
	Rating              int       `json:"rating"`
	ArticleId           int       `json:"article_id"`
	CommentForCommentId string    `json:"comment_for_comment_id"`
	Publisher           Publisher `json:"publisher"`
}

type Publisher struct {
	Username string `json:"username"`
	Login    string `json:"login"`
}
