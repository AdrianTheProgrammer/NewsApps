package handlers

import (
	"newsapps/internal/features/comments"

	"github.com/labstack/echo/v4"
)

type CommentsHand struct {
	srv comments.Services
}

func NewCommentsHand(s comments.Services) comments.Handlers {
	return &CommentsHand{
		srv: s,
	}
}

func (ch *CommentsHand) CreateComment() echo.HandlerFunc {
	// Placeholder
	return nil
}

func (ch *CommentsHand) UpdateComment() echo.HandlerFunc {
	// Placeholder
	return nil
}

func (ch *CommentsHand) DeleteComment() echo.HandlerFunc {
	// Placeholder
	return nil
}
