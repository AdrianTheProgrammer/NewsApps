package handlers

import (
	"newsapps/internal/features/comments"

	"github.com/labstack/echo/v4"
)

type CommentsHand struct {
	srv comments.CServices
}

func NewCommentsHand(s comments.CServices) comments.CHandlers {
	return &CommentsHand{
		srv: s,
	}
}

func (ch *CommentsHand) CreateComment(c echo.Context) error {
	// Placeholder
	return nil
}

func (ch *CommentsHand) UpdateComment(c echo.Context) error {
	// Placeholder
	return nil
}

func (ch *CommentsHand) DeleteComment(c echo.Context) error {
	// Placeholder
	return nil
}
