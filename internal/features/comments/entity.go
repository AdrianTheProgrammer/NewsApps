package comments

import "github.com/labstack/echo/v4"

type Comments struct {
	ID        uint
	ArticleID uint
	UserID    uint
	Content   string
}

type Handlers interface {
	CreateComment() echo.HandlerFunc
	UpdateComment() echo.HandlerFunc
	DeleteComment() echo.HandlerFunc
}

type Services interface {
}

type Query interface {
}
