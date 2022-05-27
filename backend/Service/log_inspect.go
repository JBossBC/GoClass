package Service

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"goClass/backend/Dao/Cache"
	"net/http"
)

//登录校验的中间件，检查cookie是否为有效时长内服务器所生成
//若请求中包含多个同名字cookie,那么context.cookie的处理方式具体是什么

func MiddleLogInspect(context *gin.Context) (string, error) {
	cookie, err := context.Cookie("goClass")
	if errors.Is(err, http.ErrNoCookie) {
		return "", err
	}
	userName, err := Cache.NewCookieDao().SearchCookieFromCache(cookie)
	return userName, err
}
