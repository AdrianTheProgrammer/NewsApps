package services

import (
	"newsapps/internal/features/users"
	"newsapps/internal/utils"
)

type UsersSrv struct {
	qry users.UQuery
	pu  utils.PassUtilInterface
	tu  utils.TokenUtilInterface
}

func NewUsersSrv(q users.UQuery, p utils.PassUtilInterface, t utils.TokenUtilInterface) users.UServices {
	return &UsersSrv{
		qry: q,
		pu:  p,
		tu:  t,
	}
}

func (us *UsersSrv) Login(username, password string) (users.Users, error) {
	result, err := us.qry.Login(username)
	if err != nil {
		return users.Users{}, err
	}

	err = us.pu.ComparePassword([]byte(result.Password), []byte(password))
	if err != nil {
		return users.Users{}, err
	}

	result.Token, err = us.tu.GenerateToken(result)
	if err != nil {
		return users.Users{}, err
	}

	return result, nil
}

func (us *UsersSrv) CreateUser(input users.Users) error {
	hashedPass, err := us.pu.GeneratePassword(input.Password)
	if err != nil {
		return err
	}
	input.Password = string(hashedPass)

	err = us.qry.CreateUser(input)

	return err
}

func (us *UsersSrv) ReadUser(userID uint) (users.Users, error) {
	return us.qry.ReadUser(userID)
}

func (us *UsersSrv) UpdateUser(userID uint, input users.Users) error {
	hashedPass, err := us.pu.GeneratePassword(input.Password)
	if err != nil {
		return err
	}
	input.Password = string(hashedPass)

	return us.qry.UpdateUser(userID, input)
}

func (us *UsersSrv) DeleteUser(userID uint) error {
	return us.qry.DeleteUser(userID)
}
