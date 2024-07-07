package services

import "newsapps/internal/features/users"

type UsersSrv struct {
	qry users.UQuery
}

func NewUsersSrv(q users.UQuery) users.UServices {
	return &UsersSrv{
		qry: q,
	}
}
