package repositories

import (
	"newsapps/internal/features/articles"
	c_rep "newsapps/internal/features/comments/repositories"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	UserID      uint
	Title       string
	Content     string
	ImageSource string
	Comments    []c_rep.Comments `gorm:"foreignKey:ArticleID"`
}

func (a *Article) toArticleEntity() articles.Article {
	return articles.Article{
		ID:          a.ID,
		UserID:      a.UserID,
		Title:       a.Title,
		Content:     a.Content,
		ImageSource: a.ImageSource,
	}
}

func toArticleData(input articles.Article) Article {
	return Article{
		UserID:      input.UserID,
		Title:       input.Title,
		Content:     input.Content,
		ImageSource: input.ImageSource,
	}
}
