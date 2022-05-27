package Controller

import (
	"goClass/backend/Repository"
	"goClass/backend/Service"
	gorm2 "gorm.io/gorm"
	"log"
	"strconv"
)

func AddArticle(header string, context string, userName string) Service.StatusCode {
	article := Repository.Article{
		Model:    gorm2.Model{},
		UserName: userName,
		Header:   header,
		Context:  context,
	}
	err := Service.NewArticleServer().AddArticle(article)
	if err != nil {
		log.Println(err)
		return 400
	}
	return 200
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

func FindArticle(TargetUserName string, number int) (*Service.PageArticle, error) {
	articleList, err := Service.NewArticleServer().FindArticle(TargetUserName, number)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	article := &Service.PageArticle{
		ArticleList: articleList,
	}
	return article, nil
}
