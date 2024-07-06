package repositories

import (
	"gorm.io/gorm"
)

type Comments struct {
	gorm.Model
	ArticleID uint
	UserID    uint
	Content   string
}
