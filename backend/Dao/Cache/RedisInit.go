package Cache

import (
	red "github.com/garyburd/redigo/redis"
	"log"
	"strings"
	"time"
)

type RedisConfig struct {
	Hostname string `mapstructure:"hostname"`
	Port     int    `mapstructure:"port"`
}

var (
	redisPool *red.Pool
)

func InitRedis(RedisInfo RedisConfig) {
	//redis init
	RedisInfoBuilder := strings.Builder{}
	RedisInfoBuilder.WriteString(RedisInfo.Hostname)
	RedisInfoBuilder.WriteString(":")
	RedisInfoBuilder.WriteString(string(RedisInfo.Port))
	RedisPool := &red.Pool{
		MaxIdle:     256,
		MaxActive:   0,
		IdleTimeout: time.Duration(120),
		Dial: func() (red.Conn, error) {
			return red.Dial(
				"tcp",
				"127.0.0.1:6379",
				red.DialReadTimeout(time.Duration(1000)*time.Millisecond),
				red.DialWriteTimeout(time.Duration(1000)*time.Millisecond),
				red.DialConnectTimeout(time.Duration(1000)*time.Millisecond),
				red.DialDatabase(0),
				//red.DialPassword(""),
			)
		},
	}
	redisPool = RedisPool
	log.Println("redis configuration init is ok")
}
func GetRedisConnection() red.Conn {
	return redisPool.Get()
}
