package comments

import "github.com/labstack/echo/v4"

type Comments struct {
	ID        uint
	ArticleID uint
	UserID    uint
	Content   string
}

type CHandlers interface {
	CreateComment(echo.Context) error
	UpdateComment(echo.Context) error
	DeleteComment(echo.Context) error
}

type CServices interface {
}

type CQuery interface {
}
