package repositories

import (
	"newsapps/internal/features/articles"
	c_enty "newsapps/internal/features/comments"
	c_repo "newsapps/internal/features/comments/repositories"

	"gorm.io/gorm"
)

type ArticlesQueries struct {
	db *gorm.DB
}

func NewArticlesQueries(connection *gorm.DB) articles.AQueries {
	return &ArticlesQueries{
		db: connection,
	}
}

func (aq *ArticlesQueries) ShowAllArticles() ([]articles.Article, error) {
	var result []Article
	var resultConvert []articles.Article
	err := aq.db.Find(&result).Error
	if err != nil {
		return []articles.Article{}, err
	}
	for _, v := range result {
		resultConvert = append(resultConvert, v.toArticleEntity())
	}
	return resultConvert, nil
}
func (aq *ArticlesQueries) ReadArticle(ID uint) (articles.Article, error) {
	var result Article
	err := aq.db.First(&result, ID).Error
	if err != nil {
		return articles.Article{}, err
	}
	return result.toArticleEntity(), nil
}
func (aq *ArticlesQueries) CreateArticle(newArticle articles.Article) error {
	cnv := toArticleData(newArticle)
	err := aq.db.Create(&cnv).Error
	if err != nil {
		return err
	}
	return nil
}
func (aq *ArticlesQueries) UpdateArticle(updatedArticle articles.Article, articleID uint) error {
	cnv := toArticleData(updatedArticle)
	cnv.ID = articleID
	err := aq.db.Model(&cnv).Updates(Article{Content: cnv.Content, Title: cnv.Title, ImageSource: cnv.ImageSource}).Error
	if err != nil {
		return err
	}
	return nil
}
func (aq *ArticlesQueries) DeleteArticle(ID uint) error {
	var result Article
	err := aq.db.Delete(&result, ID).Error
	if err != nil {
		return err
	}
	return nil
}

func (aq *ArticlesQueries) ShowAllComments(articleID uint) ([]c_enty.Comment, error) {
	var result []c_repo.Comment
	var resultConvert []c_enty.Comment
	err := aq.db.Where("article_id = ?", articleID).Find(&result).Error
	if err != nil {
		return []c_enty.Comment{}, err
	}
	for _, v := range result {
		resultConvert = append(resultConvert, toCommentEntity(v))
	}
	return resultConvert, nil
}
