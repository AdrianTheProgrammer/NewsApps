package repositories

import (
	"newsapps/internal/features/articles"

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
func (aq *ArticlesQueries) UpdateArticle(updatedArticle articles.Article) error {
	cnv := toArticleData(updatedArticle)
	err := aq.db.Save(&cnv).Error
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
