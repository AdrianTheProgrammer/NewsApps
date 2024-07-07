package repositories

import (
	"newsapps/internal/features/comments"

	"gorm.io/gorm"
)

type CommentsQry struct {
	db *gorm.DB
}

func NewCommentsQry(connection *gorm.DB) comments.CQuery {
	return &CommentsQry{
		db: connection,
	}
}
