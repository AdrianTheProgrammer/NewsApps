package handlers

import (
	"newsapps/internal/features/articles"

	"github.com/labstack/echo/v4"
)

type ArticlesHand struct {
	srv articles.Services
}

func NewArticlesHand(s articles.Services) articles.Handlers {
	return &ArticlesHand{
		srv: s,
	}
}

func (ah *ArticlesHand) CreateArticle() echo.HandlerFunc {
	// Placeholder
	return nil
}
func (ah *ArticlesHand) DeleteArticle() echo.HandlerFunc {
	// Placeholder
	return nil
}
func (ah *ArticlesHand) ReadArticle() echo.HandlerFunc {
	// Placeholder
	return nil
}
func (ah *ArticlesHand) ShowAllArticles() echo.HandlerFunc {
	// Placeholder
	return nil
}

func (ah *ArticlesHand) UpdateArticle() echo.HandlerFunc {
	// Placeholder
	return nil
}
