package handlers

import (
	"newsapps/internal/features/articles"

	"github.com/labstack/echo/v4"
)

type ArticlesHand struct {
	srv articles.AServices
}

func NewArticlesHand(s articles.AServices) articles.AHandlers {
	return &ArticlesHand{
		srv: s,
	}
}

func (ah *ArticlesHand) ReadAllArticles(c echo.Context) error {
	// Placeholder
	return nil
}

func (ah *ArticlesHand) CreateArticle(c echo.Context) error {
	// Placeholder
	return nil
}

func (ah *ArticlesHand) UpdateArticle(c echo.Context) error {
	// Placeholder
	return nil
}

func (ah *ArticlesHand) DeleteArticle(c echo.Context) error {
	// Placeholder
	return nil
}
