package Dao

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"goClass/backend/Repository"
	"log"
	"strings"
	"sync"
)

type UserDao struct {
}

var (
	userDao *UserDao
	once    sync.Once
)

func NewUserDao() *UserDao {
	once.Do(func() {
		GetMysqlConnection().AutoMigrate(&Repository.User{})
		userDao = &UserDao{}
	})
	return userDao
}

//检查用户名是否存在

func (userDao *UserDao) UsernameIsExist(username string) (bool, error) {
	connection := GetMysqlConnection()
	if err := connection.Where("username = ?", username).Select("username").Find(&Repository.User{}).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		} else {
			return false, nil
		}
	}
	return true, nil
}

//保存用户数据到mysql

func (userDao *UserDao) KeepUserToDataSource(user *Repository.User) error {
	connection := GetMysqlConnection()
	if err := connection.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (userDao *UserDao) UserIsExist(user *Repository.User) bool {
	connection := GetMysqlConnection()
	var DataUser *Repository.User = &Repository.User{}
	//println(connection.Model(&Repository.User{}).Where("username = ? and password = ?", user.Username, user.Password).First(user).Error.Error())
	//if DataUser != nil {
	//	return true
	//}
	//return false
	if err := connection.Model(&Repository.User{}).Where("username = ? and password = ?", user.Username, user.Password).First(DataUser).Error; err != nil {
		if strings.Compare(err.Error(), "record not found") != 0 {
			log.Println(err.Error())
			return false
		} else {
			return false
		}
	}
	return true
}
