package handlers

import (
	"newsapps/internal/features/users"
	"newsapps/internal/helpers"
	"newsapps/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UsersHand struct {
	srv users.UServices
	tu  utils.TokenUtilInterface
}

func NewUsersHand(s users.UServices, t utils.TokenUtilInterface) users.UHandlers {
	return &UsersHand{
		srv: s,
		tu:  t,
	}
}

func (uh *UsersHand) Login(c echo.Context) error {
	// Route: /login
	var input users.Users

	err := c.Bind(&input)
	if err != nil {
		return c.JSON(400, helpers.ResponseFormat(400, "Input Error!", nil))
	}

	result, err := uh.srv.Login(input.Username, input.Password)
	if err != nil {
		return c.JSON(404, helpers.ResponseFormat(404, "User Not Found!", nil))
	}

	return c.JSON(200, helpers.ResponseFormat(200, "Login Success!", result))
}

func (uh *UsersHand) CreateUser(c echo.Context) error {
	// Route: /register
	var input users.Users

	err := c.Bind(&input)
	if err != nil {
		return c.JSON(400, helpers.ResponseFormat(400, "Input Error!", nil))
	}

	err = uh.srv.CreateUser(input)
	if err != nil {
		return c.JSON(500, helpers.ResponseFormat(500, "Server Error!", nil))
	}

	return c.JSON(201, helpers.ResponseFormat(201, "Data Inserted Successfully!", nil))
}

func (uh *UsersHand) ReadUser(c echo.Context) error {
	// Route: /users/settings
	LoginData := utils.NewTokenUtil().DecodeToken(c.Get("user").(*jwt.Token))
	result, err := uh.srv.ReadUser(LoginData.ID)
	if err != nil {
		return c.JSON(500, helpers.ResponseFormat(500, "Server Error!", nil))
	}

	return c.JSON(200, helpers.ResponseFormat(200, "User data successfully retrieved!", result))
}

func (uh *UsersHand) UpdateUser(c echo.Context) error {
	// Route: /users/edit
	var input users.Users

	err := c.Bind(&input)
	if err != nil {
		return c.JSON(400, helpers.ResponseFormat(400, "Input Error!", nil))
	}

	LoginData := uh.tu.DecodeToken(c.Get("user").(*jwt.Token))
	err = uh.srv.UpdateUser(LoginData.ID, input)
	if err != nil {
		return c.JSON(500, helpers.ResponseFormat(500, "Server Error!", nil))
	}

	return c.JSON(200, helpers.ResponseFormat(200, "User data successfully edited!", nil))
}

func (uh *UsersHand) DeleteUser(c echo.Context) error {
	// Route: /users/deactivate
	LoginData := uh.tu.DecodeToken(c.Get("user").(*jwt.Token))
	err := uh.srv.DeleteUser(LoginData.ID)
	if err != nil {
		return c.JSON(500, helpers.ResponseFormat(500, "Server Error!", nil))
	}

	return c.JSON(200, helpers.ResponseFormat(200, "User data successfully deactivated!", nil))
}
