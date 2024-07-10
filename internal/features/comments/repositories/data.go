package repositories

import (
	"newsapps/internal/features/comments"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ArticleID uint
	UserID    uint
	Content   string
}

func (a *Comment) toCommentEntity() comments.Comment {
	return comments.Comment{
		ID:        a.ID,
		ArticleID: a.ArticleID,
		UserID:    a.UserID,
		Content:   a.Content,
	}
}

func toCommentData(input comments.Comment) Comment {
	return Comment{
		ArticleID: input.ArticleID,
		UserID:    input.UserID,
		Content:   input.Content,
	}
}
