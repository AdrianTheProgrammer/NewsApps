package articles

import "github.com/labstack/echo/v4"

type Articles struct {
	ID          uint
	UserID      uint
	Title       string
	Content     string
	ImageSource string
}

type AHandlers interface {
	ReadAllArticles(echo.Context) error
	CreateArticle(echo.Context) error
	UpdateArticle(echo.Context) error
	DeleteArticle(echo.Context) error
}

type AServices interface {
}

type AQuery interface {
}
