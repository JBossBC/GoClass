package Dao

import (
	"goClass/backend/Repository"
	"sync"
)

type ArticleDao struct {
}

var (
	articleDao     *ArticleDao
	articleDaoOnce sync.Once
)

func NewArticleDao() *ArticleDao {
	articleDaoOnce.Do(func() {
		articleDao = &ArticleDao{}
	})
	return articleDao
}

func (articleDao *ArticleDao) AddArticle(article *Repository.Article) error {
	connection := GetMysqlConnection()
	if err := connection.Create(article).Error; err != nil {
		return err
	}
	return nil
}
func (articleDao *ArticleDao) DeleteArticle(articleId int) error {
	connection := GetMysqlConnection()
	err := connection.Unscoped().Delete(&Repository.Article{}, articleId).Error
	return err
}
func (articleDao *ArticleDao) AccordingIDFindUserName(articleId int) (string, error) {
	connection := GetMysqlConnection()
	article := Repository.Article{}
	if err := connection.Find(&article, articleId).Error; err != nil {
		return "", err
	}
	return article.UserName, nil
}
func (articleDao *ArticleDao) FindArticleAccodingUserName(userName string, number int) ([]*Repository.Article, error) {
	articleList := make([]*Repository.Article, number)
	connection := GetMysqlConnection()
	if err := connection.Where("user_name=?", userName).Find(&articleList).Error; err != nil {
		return nil, err
	}
	return articleList, nil
}
