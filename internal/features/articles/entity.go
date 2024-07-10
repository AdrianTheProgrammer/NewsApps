package articles

import (
	"github.com/labstack/echo/v4"
)

type Article struct {
	ID          uint
	UserID      uint
	Title       string
	Content     string
	ImageSource string
}

type AHandlers interface {
	ShowAllArticles() echo.HandlerFunc
	ReadArticle() echo.HandlerFunc
	CreateArticle() echo.HandlerFunc
	UpdateArticle() echo.HandlerFunc
	DeleteArticle() echo.HandlerFunc
}

type AServices interface {
	ShowAllArticles() ([]Article, error)
	ReadArticle(ID uint) (Article, error)
	CreateArticle(newArticle Article, imgURL string, userID uint) error
	UpdateArticle(updatedArticle Article, imgURL string, userID uint, articleID uint) error
	DeleteArticle(ID uint, userID uint) error
}

type AQueries interface {
	ShowAllArticles() ([]Article, error)
	ReadArticle(ID uint) (Article, error)
	CreateArticle(newArticle Article) error
	UpdateArticle(updatedArticle Article) error
	DeleteArticle(ID uint) error
}
