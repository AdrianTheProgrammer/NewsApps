package repositories

import (
	"newsapps/internal/features/articles"

	"gorm.io/gorm"
)

type ArticlesQry struct {
	db *gorm.DB
}

func NewArticlesQry(connection *gorm.DB) articles.AQuery {
	return &ArticlesQry{
		db: connection,
	}
}
