package services

import (
	"errors"
	"newsapps/internal/features/comments"

	"gorm.io/gorm"
)

type commentServices struct {
	qry comments.CQueries
}

func NewCommentsServices(q comments.CQueries) comments.CServices {
	return &commentServices{
		qry: q,
	}
}

func (cs *commentServices) ShowAllComments(articleID uint) ([]comments.Comment, error) {
	result, err := cs.qry.ShowAllComments(articleID)
	msg := "terjadi kesalahan pada server"
	if err != nil {
		//log.Println("Show All Articles sql error:", err.Error())
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			msg = "data tidak ditemukan"
		}
		return []comments.Comment{}, errors.New(msg)
	}
	return result, nil
}
func (cs *commentServices) ReadComment(ID uint, articleID uint) (comments.Comment, error) {
	result, err := cs.qry.ReadComment(ID)
	msg := "terjadi kesalahan pada server"
	if err != nil {
		//log.Println("Show All Articles sql error:", err.Error())
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			msg = "data tidak ditemukan"
		}
		return comments.Comment{}, errors.New(msg)
	}
	if result.ArticleID != articleID {
		msg = "data tidak ditemukan"
		return comments.Comment{}, errors.New(msg)
	}
	return result, nil
}
func (cs *commentServices) CreateComment(newComment comments.Comment, userID uint, articleID uint) error {
	msg := "terjadi kesalahan pada server"
	newComment.ArticleID = articleID
	newComment.UserID = userID
	err := cs.qry.CreateComment(newComment)
	if err != nil {
		return errors.New(msg)
	}
	return nil
}
func (cs *commentServices) UpdateComment(updatedComment comments.Comment, userID uint, commentID uint) error {
	msg := "terjadi kesalahan pada server"

	result, err := cs.qry.ReadComment(commentID)
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

	err = cs.qry.UpdateComment(updatedComment, commentID)
	if err != nil {
		return errors.New(msg)
	}

	return nil
}
func (cs *commentServices) DeleteComment(ID uint, userID uint) error {
	msg := "terjadi kesalahan pada server"

	result, err := cs.qry.ReadComment(ID)
	if err != nil {
		//log.Println("Show All Articles sql error:", err.Error())
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			msg = "data comment tidak ditemukan"
		}
		return errors.New(msg)
	}
	if result.UserID != userID {
		msg = "data user tidak sesuai"
		return errors.New(msg)
	}

	err = cs.qry.DeleteComment(ID)
	if err != nil {
		//log.Println("Show All Articles sql error:", err.Error())
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			msg = "data tidak ditemukan"
		}
		return errors.New(msg)
	}
	return nil
}
