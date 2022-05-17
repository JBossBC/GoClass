package Dao

import (
	"goClass/backend/Repository"
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
		userDao = &UserDao{}
	})
	return userDao
}

//检查用户名是否存在

func (userDao *UserDao) UsernameIsExist(username string) (bool, error) {
	connection := GetMysqlConnection()
	if err := connection.Where("username = ?", username).Select("username").Find(&Repository.User{}).Error; err != nil {
		if errInt := strings.Compare(err.Error(), "record not found"); errInt != 0 {
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
