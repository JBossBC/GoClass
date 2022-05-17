package Dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"goClass/backend/Repository"
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
	if err := connection.Where("username = ? and password = ?", user.Username, user.Password).Find(&user).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println(err.Error())
			return false
		} else {
			return false
		}
	}
	return true
}
