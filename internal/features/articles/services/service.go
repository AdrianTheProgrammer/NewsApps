package services

import (
	"errors"
	"newsapps/internal/features/articles"

	"gorm.io/gorm"
)

type articleServices struct {
	qry articles.AQueries
}

func NewArticlesServices(q articles.AQueries) articles.AServices {
	return &articleServices{
		qry: q,
	}
}

func (as *articleServices) ShowAllArticles() ([]articles.Article, error) {
	result, err := as.qry.ShowAllArticles()
	msg := "terjadi kesalahan pada server"
	if err != nil {
		//log.Println("Show All Articles sql error:", err.Error())
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			msg = "data tidak ditemukan"
		}
		return []articles.Article{}, errors.New(msg)
	}

	for x, v := range result {
		result[x].Comments, err = as.qry.ShowAllComments(v.ID)
		if err != nil {
			msg = "data comment article tidak ditemukan"
		}
	}
	return result, nil
}
func (as *articleServices) ReadArticle(ID uint) (articles.Article, error) {
	result, err := as.qry.ReadArticle(ID)
	msg := "terjadi kesalahan pada server"
	if err != nil {
		//log.Println("Show All Articles sql error:", err.Error())
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			msg = "data tidak ditemukan"
		}
		return articles.Article{}, errors.New(msg)
	}
	result.Comments, err = as.qry.ShowAllComments(result.ID)
	if err != nil {
		msg = "data comment article tidak ditemukan"
	}
	return result, nil
}
func (as *articleServices) CreateArticle(newArticle articles.Article, imgURL string, userID uint) error {
	msg := "terjadi kesalahan pada server"
	newArticle.ImageSource = imgURL
	newArticle.UserID = userID
	err := as.qry.CreateArticle(newArticle)
	if err != nil {
		return errors.New(msg)
	}
	return nil
}
func (as *articleServices) UpdateArticle(updatedArticle articles.Article, imgURL string, userID uint, articleID uint) error {
	msg := "terjadi kesalahan pada server"

	result, err := as.qry.ReadArticle(articleID)
	if err != nil {
		//log.Println("Show All Articles sql error:", err.Error())
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			msg = "data article tidak ditemukan"
		}
		return errors.New(msg)
	}

	if result.UserID != userID {
		msg = "data user tidak ditemukan"
		return errors.New(msg)
	}
	updatedArticle.ID = articleID
	updatedArticle.ImageSource = imgURL
	err = as.qry.UpdateArticle(updatedArticle, articleID)
	if err != nil {
		return errors.New(msg)
	}

	return nil
}
func (as *articleServices) DeleteArticle(ID uint, userID uint) error {
	msg := "terjadi kesalahan pada server"

	result, err := as.qry.ReadArticle(ID)
	if err != nil {
		//log.Println("Show All Articles sql error:", err.Error())
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			msg = "data article tidak ditemukan"
		}
		return errors.New(msg)
	}
	if result.UserID != userID {
		msg = "data user tidak sesuai"
		return errors.New(msg)
	}

	err = as.qry.DeleteArticle(ID)
	if err != nil {
		//log.Println("Show All Articles sql error:", err.Error())
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			msg = "data tidak ditemukan"
		}
		return errors.New(msg)
	}
	return nil
}
