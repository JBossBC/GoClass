package Cache

import (
	"fmt"
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
func (cookieDao *cookie) KeepCookieToCache(cookie string) {
	connection := GetRedisConnection()
	_, err := connection.Do("Set", cookie, true)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = connection.Do("expire", cookie, "10000")
	if err != nil {
		fmt.Println(err.Error())
	}

}
