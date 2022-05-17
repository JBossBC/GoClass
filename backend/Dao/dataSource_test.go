package Dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"goClass/backend/Repository"
	"goClass/backend/util"
	"testing"
)

//test dataConfiguration
func TestDataSourceInit(t *testing.T) {
	connection := GetMysqlConnection()
	//newUser := Repository.User{}
	//connection.CreateTable(&Repository.User{})
	user := Repository.User{
		Model:    gorm.Model{},
		Username: "1577002722",
		Password: util.MD5EnCrypto("jiang19780809"),
	}
	connection.Create(&user)
	fmt.Println(user.Password)
	//connection.Where("username= ? ", user.Username).Select("username").Find(&Repository.User{}).Scan(&newUser)
	//connection.Find(&Repository.User{}).Scan(&newUser)
	//if err := connection.Where("username = 15770027222").Select("username").Find(&Repository.User{}).Error; err != nil {
	//	//fmt.Println(err.type())
	//	fmt.Printf(err.Error())
	//	fmt.Println(" have error")
	//}
	//row.Scan(&newUser)
	//rows, err := db.Rows()
	//fmt.Println(newUser)
}

//检查 userIsExist function 是否正确
func TestUserIsExist(t *testing.T) {
	exist, err := NewUserDao().UsernameIsExist("15770027222")
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println(exist)
}

func TestKeepUser(t *testing.T) {
	user := Repository.User{
		Model:    gorm.Model{},
		Username: "135491849165",
		Password: util.MD5EnCrypto("45984561894"),
	}
	err := NewUserDao().KeepUserToDataSource(&user)
	if err != nil {
		return
	}
}
