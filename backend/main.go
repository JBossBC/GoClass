package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"goClass/backend/Controller"
	"goClass/backend/Dao/Cache"
	"goClass/backend/Service"
	"log"
	"net/http"
	"strconv"
)

func main() {
	engine := gin.Default()
	initRoute(engine)
	engine.Run()
}

const DefaultNumber = 5

//请求接受过程中包含中文会解析出现乱码
func initRoute(engine *gin.Engine) {
	engine.Handle(http.MethodPost, "/register", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		rePassword := context.PostForm("rePassword")
		//controller层处理username和password
		//保存数据到数据库,重定向成功页面
		statusCode := Controller.Register(username, password, rePassword)
		context.JSON(int(statusCode), "congratulations you register success")
	})
	engine.Handle(http.MethodGet, "/log", func(context *gin.Context) {
		username := context.Query("username")
		password := context.Query("password")
		isTrue := Controller.Log(username, password)
		//保存会话信息，便于后续功能的利用,value使用UUID生成
		if isTrue {
			tempUUID := uuid.New()
			Cache.NewCookieDao().KeepCookieToCache(tempUUID, username)
			context.SetCookie("goClass", tempUUID, 0, "/", ".", true, true)
			context.JSON(http.StatusOK, "congratulations you log success")
		}
	})
	//addArticle route  post request params:header context cookie:log_auth
	engine.Handle(http.MethodPost, "/addArticle", func(context *gin.Context) {
		userName, err := Service.MiddleLogInspect(context)
		if err != nil {
			context.JSON(300, "You haven't signed in yet,please retry after log in  ")
		}
		tempHeader := context.PostForm("header")
		header := strconv.QuoteToASCII(tempHeader)
		ArticleContext := context.PostForm("context")
		article := Controller.AddArticle(header, ArticleContext, userName)
		if article != 200 {
			log.Println(err)
			context.JSON(int(article), err.Error())
		}
		context.JSON(200, "congratulation you add article successful")
	})
	//deleteArticleRoute   params:article_id cookie:log_auth
	engine.Handle(http.MethodGet, "/deleteArticle", func(context *gin.Context) {
		articleID := context.Query("articleID")
		logUserName, err := Service.MiddleLogInspect(context)
		if err != nil {
			context.JSON(400, err)
		}
		statusCode := Controller.DeleteArticle(articleID, logUserName)
		if statusCode == 200 {
			context.JSON(int(statusCode), "congratulation you delete article successful")
		}
	})
	//GET article params:targetUserName number
	engine.Handle(http.MethodGet, "getArticle", func(context *gin.Context) {
		var number int
		targetUserName := context.Query("targetUserName")
		if targetUserName == "" {
			context.JSON(300, "send target UserName is empty,please send you need to find the username of article")
		}
		numberStr := context.Query("number")
		if numberStr == "" {
			number = DefaultNumber
		} else {
			number, _ = strconv.Atoi(numberStr)
		}
		article, err := Controller.FindArticle(targetUserName, number)
		if err != nil {
			context.JSON(400, err)
		}
		context.JSON(200, article)
	})
}
