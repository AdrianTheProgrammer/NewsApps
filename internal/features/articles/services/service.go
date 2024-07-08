package services

import (
	"newsapps/internal/features/articles"
)

type articleServices struct {
	qry articles.Queries
}

func NewArticlesServices(q articles.Queries) articles.Services {
	return &articleServices{
		qry: q,
	}
}

func (as *articleServices) ShowAllArticles() ([]articles.Article, error) {
	// placeholder
	return nil, nil
}
func (as *articleServices) ReadArticle(ID uint) (articles.Article, error) {
	// placeholder
	return articles.Article{}, nil
}
func (as *articleServices) CreateArticle(newArticle articles.Article) error {
	// placeholder
	return nil
}
func (as *articleServices) UpdateArticle(updatedArticle articles.Article) error {
	// placeholder
	return nil
}
func (as *articleServices) DeleteArticle(ID uint) error {
	// placeholder
	return nil
}
