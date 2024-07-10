package handlers

import (
	"newsapps/internal/features/comments"
)

type NewCommentsRequest struct {
	Content string `json:"content" form:"content"`
}

func ToCommentEntity(a NewCommentsRequest) comments.Comment {
	return comments.Comment{
		Content: a.Content,
	}
}
