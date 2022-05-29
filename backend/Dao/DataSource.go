package Dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"goClass/backend/Dao/Cache"
	"log"
	"os"
	"strings"
)

type Configure struct {
	DataInfo  string            `mapstructure:"name"`
	MysqlInfo MysqlConfig       `mapstructure:"mysql"`
	RedisInfo Cache.RedisConfig `mapstructure:"redis"`
}

func init() {
	v := viper.New()
	v.SetConfigName("dataSource")
	configureFile := strings.Builder{}
	wd, _ := os.Getwd()
	configureFile.WriteString(wd)

	configureFile.WriteString("/backend/Dao/configure.yaml")
	v.SetConfigFile(configureFile.String())

	//test DataSource configure
	//v.SetConfigFile("./configure.yaml")

	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic("dataSource Configure error")
	}
	dataSource := Configure{}
	err = v.Unmarshal(&dataSource)
	if err != nil {
		panic("viper Unmarshal error")
	}
	log.Println("dataSource configuration init success")
	InitMysql(dataSource)
	Cache.InitRedis(dataSource.RedisInfo)
}
