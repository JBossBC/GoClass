package Controller

import (
	"goClass/backend/Repository"
	"goClass/backend/Service"
	gorm2 "gorm.io/gorm"
	"log"
	"strconv"
)

func AddArticle(header string, context string, userName string) (Service.StatusCode, int) {
	article := Repository.Article{
		Model:    gorm2.Model{},
		UserName: userName,
		Header:   header,
		Context:  context,
	}
	nowId, err := Service.NewArticleServer().AddArticle(article)
	if err != nil {
		log.Println(err)
		return 400, 0
	}
	return 200, nowId
}

func DeleteArticle(articleIdString string, userName string) Service.StatusCode {
	articleId, _ := strconv.ParseInt(articleIdString, 10, 32)
	err := Service.NewArticleServer().DeleteArticle(int(articleId), userName)
	if err != nil {
		log.Println(err)
		return 400
	}
	return 200
}

func FindArticle(TargetUserName string, number int) (*Service.ArticlePage, error) {
	articleList, err := Service.NewArticleServer().FindArticle(TargetUserName, number)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return articleList, nil
}
func UpdateArticle(header string, context string, updateId int) error {
	article := &Repository.Article{
		Model:    gorm2.Model{ID: uint(updateId)},
		UserName: "",
		Header:   header,
		Context:  context,
	}
	return Service.NewArticleServer().UpdateArticle(article)
}
