package Dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	MysqlInfoBuilder.WriteString("?charset=utf8&parseTime=True&loc=Local")
	//fmt.Println(MysqlInfoBuilder.String())
	//gorm 1.0版本链接mysql的方式,但结合实际开发来说，此时我们可以利用预编译语句和默认跳过事务来提高效率，所以我们采取gorm2.0的方式，自定义gorm.config
	//mysqlConnection, err = gorm.Open("mysql", MysqlInfoBuilder.String(),&gorm.)
	mysqlConnection, err = gorm.Open(mysql.Open(MysqlInfoBuilder.String()), &gorm.Config{PrepareStmt: true, SkipDefaultTransaction: true})
	if err != nil {
		panic("dataSource connection error " + err.Error())
	}
	log.Println("mysql connection init is ok")
}
func GetMysqlConnection() *gorm.DB {
	return mysqlConnection
}
