package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"goClass/backend/Controller"
	_ "goClass/backend/Dao"
	"goClass/backend/Dao/Cache"
	"goClass/backend/Service"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	engine := gin.Default()
	openLog(engine)
	initRoute(engine)
	engine.Run()
}
func openLog(engine *gin.Engine) {
	file, err := os.OpenFile("C:\\ProgramData\\goClass\\goClass.log", os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	gin.DisableConsoleColor()
	config := gin.LoggerConfig{Output: file}
	gin.LoggerWithConfig(config)
	gin.DefaultWriter = io.MultiWriter(file)
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
		} else {
			context.JSON(300, "log failed,maybe your password or username is wrong")
		}
	})
	//addArticle route  post request params:header context cookie:log_auth
	engine.Handle(http.MethodPost, "/addArticle", func(context *gin.Context) {
		userName, err := Service.MiddleLogInspect(context)
		if err != nil || userName == "" {
			context.JSON(300, "You haven't signed in yet,please retry after log in  ")
			return
		}
		Header := context.PostForm("header")
		ArticleContext := context.PostForm("context")
		article, nowId := Controller.AddArticle(Header, ArticleContext, userName)
		if article != 200 {
			log.Println(err)
			context.JSON(int(article), err.Error())
			return
		}
		context.JSON(200, "congratulation you add article successful articleId is "+strconv.Itoa(nowId))
	})
	//deleteArticleRoute   params:articleID cookie:log_auth
	engine.Handle(http.MethodGet, "/deleteArticle", func(context *gin.Context) {
		articleID := context.Query("articleID")
		logUserName, err := Service.MiddleLogInspect(context)
		if err != nil {
			context.JSON(400, err)
			return
		}
		statusCode := Controller.DeleteArticle(articleID, logUserName)
		if statusCode == 200 {
			context.JSON(int(statusCode), "congratulation you delete article successful")
			return
		}
		context.JSON(int(statusCode), "delete article failed ,please wait a moment")
	})
	//GET article params:targetUserName number
	engine.Handle(http.MethodGet, "/getArticle", func(context *gin.Context) {
		var number int
		targetUserName := context.Query("targetUserName")
		if targetUserName == "" {
			context.JSON(300, "send target UserName is empty,please send you need to find the username of article")
			return
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
			return
		}
		context.JSON(200, article)
	})
	//updateArticle  params:header,context,updateID                LogSession
	engine.Handle(http.MethodPost, "/updateArticle", func(context *gin.Context) {
		header := context.PostForm("header")
		ArticleContext := context.PostForm("context")
		updateIdStr := context.PostForm("updateId")
		updateId, _ := strconv.Atoi(updateIdStr)
		err := Controller.UpdateArticle(header, ArticleContext, updateId)
		if err != nil {
			context.JSON(400, err)
			return
		}
		context.JSON(200, "congratulation you update successful")
	})
	engine.Handle(http.MethodPost, "/uploadPicture", func(context *gin.Context) {
		file, err := context.FormFile("picture")
		if err != nil {
			context.JSON(300, "file params wrong")
			return
		}
		err = Controller.UploadPicture(file)
		if err != nil {
			context.JSON(400, "uploadFailed")
			return
		}
		context.JSON(200, "upload successful")
	})
}
