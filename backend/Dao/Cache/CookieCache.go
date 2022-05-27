package Cache

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"sync"
)

type cookie struct {
}

var (
	cookieDao  *cookie
	cookieOnce sync.Once
)

func NewCookieDao() *cookie {
	cookieOnce.Do(func() {
		cookieDao = &cookie{}
	})
	return cookieDao
}

//因为是redis连接池，所以不需要在一次请求结束后就释放connection

func (cookieDao *cookie) KeepCookieToCache(cookie string, userName string) {
	connection := GetRedisConnection()
	_, err := connection.Do("Set", cookie, userName)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = connection.Do("expire", cookie, "10000")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (cookieDao *cookie) SearchCookieFromCache(cookie string) (string, error) {
	connection := GetRedisConnection()
	reply, err := redis.String(connection.Do("get", cookie))
	if err != nil {
		return "", err
	}
	return reply, nil
}
