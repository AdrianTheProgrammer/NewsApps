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
