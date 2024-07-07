package users

import "github.com/labstack/echo/v4"

type Users struct {
	ID       uint
	Username string
	Password string
	Fullname string
	Email    string
	Phone    string
}

type UHandlers interface {
	Login(echo.Context) error
	CreateUser(echo.Context) error
	ReadUser(echo.Context) error
	UpdateUser(echo.Context) error
	DeleteUser(echo.Context) error
}

type UServices interface {
}

type UQuery interface {
}
