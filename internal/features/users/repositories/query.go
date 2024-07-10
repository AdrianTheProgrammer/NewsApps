package repositories

import (
	"newsapps/internal/features/users"

	"gorm.io/gorm"
)

type UsersQry struct {
	db *gorm.DB
}

func NewUsersQry(connection *gorm.DB) users.UQuery {
	return &UsersQry{
		db: connection,
	}
}

func (uq *UsersQry) Login(username string) (users.Users, error) {
	var result Users
	err := uq.db.Where("username = ?", username).First(&result).Error
	rescnv := ToUsersEntity(result)

	if err != nil {
		return users.Users{}, err
	}

	return rescnv, nil
}

func (uq *UsersQry) CreateUser(input users.Users) error {
	cnv := ToUsersData(input)
	err := uq.db.Create(&cnv).Error
	if err != nil {
		return err
	}

	return nil
}

func (uq *UsersQry) ReadUser(userID uint) (users.Users, error) {
	var result Users
	err := uq.db.Where("id = ?", userID).First(&result).Error
	if err != nil {
		return users.Users{}, err
	}

	return ToUsersEntity(result), nil
}

func (uq *UsersQry) UpdateUser(userID uint, input users.Users) error {
	cnv := ToUsersData(input)
	cnv.ID = userID
	err := uq.db.Save(&cnv).Error
	if err != nil {
		return err
	}

	return nil
}

func (uq *UsersQry) DeleteUser(userID uint) error {
	err := uq.db.Delete(&Users{}, userID).Error
	if err != nil {
		return err
	}

	return nil
}
