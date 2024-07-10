package utils

import (
	"newsapps/configs"
	"newsapps/internal/features/users"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenUtilInterface interface {
	GenerateToken(users.Users) (string, error)
	DecodeToken(*jwt.Token) users.Users
}

type tokenUtil struct{}

func NewTokenUtil() TokenUtilInterface {
	return &tokenUtil{}
}

func (tu *tokenUtil) GenerateToken(LoginData users.Users) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = LoginData.ID
	claims["username"] = LoginData.Username
	claims["password"] = LoginData.Password
	claims["fullname"] = LoginData.Fullname
	claims["email"] = LoginData.Email
	claims["phone"] = LoginData.Phone
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	process := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := process.SignedString([]byte(configs.ImportPasskey()))

	if err != nil {
		return "", err
	}

	return result, nil
}

func (tu *tokenUtil) DecodeToken(token *jwt.Token) users.Users {
	claims := token.Claims.(jwt.MapClaims)

	var result users.Users
	result.ID = uint(claims["id"].(float64))
	result.Username = claims["username"].(string)
	result.Password = claims["password"].(string)
	result.Fullname = claims["fullname"].(string)
	result.Email = claims["email"].(string)
	result.Phone = claims["phone"].(string)

	return result
}
