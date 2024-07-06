package repositories

import (
	c_rep "newsapps/internal/features/comments/repositories"

	"gorm.io/gorm"
)

type Articles struct {
	gorm.Model
	UserID      uint
	Title       string
	Content     string
	ImageSource string
	Comments    c_rep.Comments `gorm:"foreignKey:ArticleID"`
}
