package Dao

import (
	"fmt"
	"goClass/backend/Repository"
	"goClass/backend/util"
	gorm2 "gorm.io/gorm"
	"testing"
)

//test dataConfiguration
func TestDataSourceInit(t *testing.T) {
	connection := GetMysqlConnection()
	//newUser := Repository.User{}
	//connection.CreateTable(&Repository.User{})
	user := Repository.User{
		Model:    gorm2.Model{},
		Username: "1577002722",
		Password: util.MD5EnCrypto("jiang19780809"),
	}
	println(connection.Create(&user).RowsAffected)
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
		Model:    gorm2.Model{},
		Username: "135491849165",
		Password: util.MD5EnCrypto("45984561894"),
	}
	err := NewUserDao().KeepUserToDataSource(&user)
	if err != nil {
		return
	}
}
func TestUpdata(t *testing.T) {
	GetMysqlConnection().AutoMigrate(&Repository.User{})
	GetMysqlConnection().AutoMigrate(&Repository.Article{})
}
func TestArticleDao_AddArticle(t *testing.T) {
	//GetMysqlConnection().Create(&Repository.Article{})
	GetMysqlConnection().AutoMigrate(&Repository.Article{}, &Repository.User{})
	//article := Repository.Article{
	//	Model:      gorm2.Model{},
	//	BelongUser: 9,
	//	Header:     "hello,world",
	//	Context:    "this is text data about adding article",
	//}
	u := Repository.User{}

	GetMysqlConnection().Find(&u, 9)
	article1 := Repository.Article{
		Model:   gorm2.Model{},
		Header:  "text",
		Context: "text",
	}
	dao := NewArticleDao()
	err := dao.AddArticle(&article1)
	if err != nil {
		return
	}
	//dao := NewArticleDao()
	//NewUserDao().KeepUserToDataSource(&Repository.User{
	//	Model:       gorm2.Model{},
	//	Username:    "xiyang",
	//	Password:    "1577002722",
	//})
	//err := dao.AddArticle(article1)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
}
func TestFindNowArticleId(t *testing.T) {
	println(NewArticleDao().NowArticleID())
}
