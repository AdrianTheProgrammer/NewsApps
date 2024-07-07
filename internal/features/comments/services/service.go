package services

import "newsapps/internal/features/comments"

type CommentsSrv struct {
	qry comments.CQuery
}

func NewCommentsSrv(q comments.CQuery) comments.CServices {
	return &CommentsSrv{
		qry: q,
	}
}
