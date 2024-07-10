package comments

import (
	"github.com/labstack/echo/v4"
)

type Comment struct {
	ID        uint
	ArticleID uint
	UserID    uint
	Content   string
}

type Handlers interface {
	ShowAllComments() echo.HandlerFunc
	ReadComment() echo.HandlerFunc
	CreateComment() echo.HandlerFunc
	UpdateComment() echo.HandlerFunc
	DeleteComment() echo.HandlerFunc
}

type Services interface {
	ShowAllComments(articleID uint) ([]Comment, error)
	ReadComment(ID uint) (Comment, error)
	CreateComment(newComment Comment, userID uint, articleID uint) error
	UpdateComment(updatedComment Comment, userID uint, commentID uint) error
	DeleteComment(ID uint, userID uint) error
}

type Queries interface {
	ShowAllComments(articleID uint) ([]Comment, error)
	ReadComment(ID uint) (Comment, error)
	CreateComment(newComment Comment) error
	UpdateComment(updatedComment Comment) error
	DeleteComment(ID uint) error
}
