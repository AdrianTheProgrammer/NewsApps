package handlers

import (
	"newsapps/internal/features/comments"
)

type NewCommentsRequest struct {
	ArticleID uint   `json:"articleID" form:"articleID"`
	Content   string `json:"content" form:"content"`
}

func ToCommentEntity(a NewCommentsRequest) comments.Comment {
	return comments.Comment{
		ArticleID: a.ArticleID,
		Content:   a.Content,
	}
}
