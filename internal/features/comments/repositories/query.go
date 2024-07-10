package repositories

import (
	"newsapps/internal/features/comments"

	"gorm.io/gorm"
)

type CommentsQueries struct {
	db *gorm.DB
}

func NewCommentsQueries(connection *gorm.DB) comments.Queries {
	return &CommentsQueries{
		db: connection,
	}
}

func (cq *CommentsQueries) ShowAllComments(articleID uint) ([]comments.Comment, error) {
	var result []Comment
	var resultConvert []comments.Comment
	err := cq.db.Where("article_id = ?", articleID).Find(&result).Error
	if err != nil {
		return []comments.Comment{}, err
	}
	for _, v := range result {
		resultConvert = append(resultConvert, v.toCommentEntity())
	}
	return resultConvert, nil
}
func (cq *CommentsQueries) ReadComment(ID uint) (comments.Comment, error) {
	var result Comment
	err := cq.db.First(&result, ID).Error
	if err != nil {
		return comments.Comment{}, err
	}
	return result.toCommentEntity(), nil
}
func (cq *CommentsQueries) CreateComment(newComment comments.Comment) error {
	cnv := toCommentData(newComment)
	err := cq.db.Create(&cnv).Error
	if err != nil {
		return err
	}
	return nil
}
func (cq *CommentsQueries) UpdateComment(updatedComment comments.Comment) error {
	cnv := toCommentData(updatedComment)
	err := cq.db.Save(&cnv).Error
	if err != nil {
		return err
	}
	return nil
}
func (cq *CommentsQueries) DeleteComment(ID uint) error {
	var result Comment
	err := cq.db.Delete(&result, ID).Error
	if err != nil {
		return err
	}
	return nil
}
