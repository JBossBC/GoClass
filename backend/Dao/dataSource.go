package Dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
	"strings"
)

var (
	mysqlConnection *gorm.DB
)

type Configure struct {
	DataInfo  string      `mapstructure:"name"`
	MysqlInfo MysqlConfig `mapstructure:"mysql"`
}
type MysqlConfig struct {
	//Host     string `mapstructure:"mysql.hostname"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`
}

func init() {
	v := viper.New()
	v.SetConfigName("dataSource")
	v.SetConfigFile("./configure.yaml")
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
	//mysql init
	//"user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	builder := strings.Builder{}
	builder.WriteString(dataSource.MysqlInfo.Name)
	builder.WriteString(":")
	builder.WriteString(dataSource.MysqlInfo.Password)
	builder.WriteString("@tcp/")
	builder.WriteString(dataSource.MysqlInfo.DBName)
	builder.WriteString("?charset=gbk&parseTime=True&loc=Local")
	fmt.Println(builder.String())
	mysqlConnection, err = gorm.Open("mysql", builder.String())
	if err != nil {
		panic("dataSource connection error " + err.Error())
	}
	log.Println("connection init is ok")
}
func GetMysqlConnection() *gorm.DB {
	return mysqlConnection
}
