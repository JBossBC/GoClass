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
func (articleServer *ArticleServer) AddArticle(article Repository.Article) (int, error) {
	err := Dao.NewArticleDao().AddArticle(&article)
	id := Dao.NewArticleDao().NowArticleID()
	return id, err
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

type ArticlePage struct {
	Article []*ShowArticle
}
type ShowArticle struct {
	Username string `json:"Username"`
	Header   string `json:"Header"`
	Context  string `json:"Context"`
}

func (articleServer *ArticleServer) FindArticle(targetUserName string, number int) (*ArticlePage, error) {
	name, err := Dao.NewArticleDao().FindArticleAccodingUserName(targetUserName, number)
	if err != nil {
		return nil, err
	}
	page := ArticlePage{Article: make([]*ShowArticle, len(name))}
	for i := 0; i < len(name); i++ {
		page.Article[i] = &ShowArticle{
			Username: name[i].UserName,
			Header:   name[i].Header,
			Context:  name[i].Context,
		}
	}
	return &page, nil
}

func (articleServer *ArticleServer) UpdateArticle(atricle *Repository.Article) error {
	return Dao.NewArticleDao().UpdateArticle(atricle)
}
