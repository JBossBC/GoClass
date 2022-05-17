package Dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"strings"
)

var (
	mysqlConnection *gorm.DB
)

type MysqlConfig struct {
	//Host     string `mapstructure:"mysql.hostname"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`
}

func InitMysql(dataSource Configure) {
	//mysql init
	//"user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	var err error
	MysqlInfoBuilder := strings.Builder{}
	MysqlInfoBuilder.WriteString(dataSource.MysqlInfo.Name)
	MysqlInfoBuilder.WriteString(":")
	MysqlInfoBuilder.WriteString(dataSource.MysqlInfo.Password)
	MysqlInfoBuilder.WriteString("@tcp/")
	MysqlInfoBuilder.WriteString(dataSource.MysqlInfo.DBName)
	MysqlInfoBuilder.WriteString("?charset=gbk&parseTime=True&loc=Local")
	fmt.Println(MysqlInfoBuilder.String())
	mysqlConnection, err = gorm.Open("mysql", MysqlInfoBuilder.String())
	if err != nil {
		panic("dataSource connection error " + err.Error())
	}
	log.Println("mysql connection init is ok")
}
func GetMysqlConnection() *gorm.DB {
	return mysqlConnection
}
