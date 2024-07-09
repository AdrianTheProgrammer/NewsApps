package services

import "newsapps/internal/features/comments"

type CommentsSrv struct {
	qry comments.Query
}

func NewCommentsSrv(q comments.Query) comments.Services {
	return &CommentsSrv{
		qry: q,
	}
}
