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

type Handlers interface {
	ShowAllArticles() echo.HandlerFunc
	ReadArticle() echo.HandlerFunc
	CreateArticle() echo.HandlerFunc
	UpdateArticle() echo.HandlerFunc
	DeleteArticle() echo.HandlerFunc
}

type Services interface {
	ShowAllArticles() ([]Article, error)
	ReadArticle(ID uint) (Article, error)
	CreateArticle(newArticle Article, imgURL string, userID uint) error
	UpdateArticle(updatedArticle Article, imgURL string, userID uint) error
	DeleteArticle(ID uint, userID uint) error
}

type Queries interface {
	ShowAllArticles() ([]Article, error)
	ReadArticle(ID uint) (Article, error)
	CreateArticle(newArticle Article) error
	UpdateArticle(updatedArticle Article) error
	DeleteArticle(ID uint) error
}
