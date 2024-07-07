package handlers

import (
	"newsapps/internal/features/users"

	"github.com/labstack/echo/v4"
)

type UsersHand struct {
	srv users.UServices
}

func NewUsersHand(s users.UServices) users.UHandlers {
	return &UsersHand{
		srv: s,
	}
}

func (uh *UsersHand) Login(c echo.Context) error {
	// Placeholder
	return nil
}

func (uh *UsersHand) CreateUser(c echo.Context) error {
	// Placeholder
	return nil
}

func (uh *UsersHand) ReadUser(c echo.Context) error {
	// Placeholder
	return nil
}

func (uh *UsersHand) UpdateUser(c echo.Context) error {
	// Placeholder
	return nil
}

func (uh *UsersHand) DeleteUser(c echo.Context) error {
	// Placeholder
	return nil
}
