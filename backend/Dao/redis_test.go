package Dao

import (
	"goClass/backend/Dao/Cache"
	"testing"
)

func TestRedis(t *testing.T) {
	Cache.NewCookieDao().KeepCookieToCache("test")
}
