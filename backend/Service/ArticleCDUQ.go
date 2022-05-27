package Service

import (
	"fmt"
	"goClass/backend/Dao"
	"goClass/backend/Repository"
	"strings"
	"sync"
)

type ArticleServer struct {
}

var (
	articleServer *ArticleServer
	articleOnce   sync.Once
)

func NewArticleServer() *ArticleServer {
	articleOnce.Do(func() {
		articleServer = &ArticleServer{}
	})
	return articleServer
}
func (articleServer *ArticleServer) AddArticle(article Repository.Article) error {
	return Dao.NewArticleDao().AddArticle(&article)
}
func (articleServer *ArticleServer) DeleteArticle(articleID int, userName string) error {
	if !articleServer.havePowerToUpdateArticle(articleID, userName) {
		return fmt.Errorf("you havent power to delete article")
	}
	err := Dao.NewArticleDao().DeleteArticle(articleID)
	if err != nil {
		return err
	}
	return nil
}
func (articleServer *ArticleServer) havePowerToUpdateArticle(articleID int, userName string) bool {
	AccordingIdObtainName, err := Dao.NewArticleDao().AccordingIDFindUserName(articleID)
	if err != nil {
		return false
	}
	//文章所有者是否为登录用户
	if strings.Compare(AccordingIdObtainName, userName) == 0 {
		return true
	} else {
		return false
	}
}

const DefaultNumber = 5

type PageArticle struct {
	ArticleList []*Repository.Article
}

func (articleServer *ArticleServer) DefaultFindArticle(targetUserName string) ([]*Repository.Article, error) {
	return articleServer.FindArticle(targetUserName, DefaultNumber)
}
func (articleServer *ArticleServer) FindArticle(targetUserName string, number int) ([]*Repository.Article, error) {
	return Dao.NewArticleDao().FindArticleAccodingUserName(targetUserName, number)
}
