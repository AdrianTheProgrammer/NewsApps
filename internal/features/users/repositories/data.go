package repositories

import (
	a_rep "newsapps/internal/features/articles/repositories"
	c_rep "newsapps/internal/features/comments/repositories"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username string
	Password string
	Fullname string
	Email    string
	Phone    string
	Articles []a_rep.Article  `gorm:"foreignKey:UserID"`
	Comments []c_rep.Comments `gorm:"foreignKey:UserID"`
}
