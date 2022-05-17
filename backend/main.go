package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"goClass/backend/Controller"
	"goClass/backend/Dao/Cache"
	"net/http"
)

func main() {
	engine := gin.Default()
	initRoute(engine)
	engine.Run()
}

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
			Cache.NewCookieDao().KeepCookieToCache(tempUUID)
			context.SetCookie("goClass", tempUUID, 0, "/", ".", true, true)
			context.JSON(http.StatusOK, "congratulations you log success")
		}
	})
}
