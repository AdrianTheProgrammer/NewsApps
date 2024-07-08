package repositories

import (
	"newsapps/internal/features/articles"

	"gorm.io/gorm"
)

type ArticlesQueries struct {
	db *gorm.DB
}

func NewArticlesQueries(connection *gorm.DB) articles.Queries {
	return &ArticlesQueries{
		db: connection,
	}
}

func (aq *ArticlesQueries) ShowAllArticles() ([]articles.Article, error) {
	// placeholder
	return nil, nil
}
func (aq *ArticlesQueries) ReadArticle(ID uint) (articles.Article, error) {
	// placeholder
	return articles.Article{}, nil
}
func (aq *ArticlesQueries) CreateArticle(newArticle articles.Article) error {
	// placeholder
	return nil
}
func (aq *ArticlesQueries) UpdateArticle(updatedArticle articles.Article) error {
	// placeholder
	return nil
}
func (aq *ArticlesQueries) DeleteArticle(ID uint) error {
	// placeholder
	return nil
}
