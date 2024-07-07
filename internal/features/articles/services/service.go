package services

import (
	"newsapps/internal/features/articles"
)

type ArticlesSrv struct {
	qry articles.AQuery
}

func NewArticlesSrv(q articles.AQuery) articles.AServices {
	return &ArticlesSrv{
		qry: q,
	}
}
