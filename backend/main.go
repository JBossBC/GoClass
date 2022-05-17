package main

import (
	"github.com/gin-gonic/gin"
	"goClass/backend/Controller"
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
		context.JSON(int(statusCode), nil)
	})
	engine.Handle(http.MethodGet, "/log", func(context *gin.Context) {

	})
}
