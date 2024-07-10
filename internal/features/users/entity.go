package users

import "github.com/labstack/echo/v4"

type Users struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Token    string `json:"token"`
}

type UHandlers interface {
	Login(echo.Context) error
	CreateUser(echo.Context) error
	ReadUser(echo.Context) error
	UpdateUser(echo.Context) error
	DeleteUser(echo.Context) error
}

type UServices interface {
	Login(string, string) (Users, error)
	CreateUser(Users) error
	ReadUser(uint) (Users, error)
	UpdateUser(uint, Users) error
	DeleteUser(uint) error
}

type UQuery interface {
	Login(string) (Users, error)
	CreateUser(Users) error
	ReadUser(uint) (Users, error)
	UpdateUser(uint, Users) error
	DeleteUser(uint) error
}
