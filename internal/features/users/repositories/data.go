package repositories

import (
	a_rep "newsapps/internal/features/articles/repositories"
	c_rep "newsapps/internal/features/comments/repositories"
	"newsapps/internal/features/users"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username string
	Password string
	Fullname string
	Email    string
	Phone    string
	Articles []a_rep.Article `gorm:"foreignKey:UserID"`
	Comments []c_rep.Comment `gorm:"foreignKey:UserID"`
}

func ToUsersData(input users.Users) Users {
	return Users{
		Username: input.Username,
		Password: input.Password,
		Fullname: input.Fullname,
		Email:    input.Email,
		Phone:    input.Phone,
	}
}

func ToUsersEntity(input Users) users.Users {
	return users.Users{
		ID:       input.ID,
		Username: input.Username,
		Password: input.Password,
		Fullname: input.Fullname,
		Email:    input.Email,
		Phone:    input.Phone,
	}
}
