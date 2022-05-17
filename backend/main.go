package main

import (
	"github.com/gin-gonic/gin"
	"goClass/backend/Controller"
	"net/http"
)

func main() {
	engine := gin.Default()
	initRoute(engine)

}

func initRoute(engine *gin.Engine) {
	engine.Handle(http.MethodPost, "/register", func(context *gin.Context) {
		username, _ := context.GetPostForm("userName")
		password, _ := context.GetPostForm("password")
		rePassword, _ := context.GetPostForm("rePassword")
		//controller层处理username和password
		//保存数据到数据库,重定向成功页面
		statusCode := Controller.Register(username, password, rePassword)
		context.JSON(int(statusCode), nil)
	})
	engine.Handle(http.MethodGet, "/log", func(context *gin.Context) {

	})
}
