package Service

import (
	"fmt"
	"goClass/backend/Controller"
	"goClass/backend/Dao"
	"goClass/backend/Repository"
	"sync"
)

const (
	UsernameExist   Controller.StatusCode = 402
	ServerError     Controller.StatusCode = 500
	SuccessRegister Controller.StatusCode = 200
)

type register struct {
}

var (
	registerService *register
	registerOnce    sync.Once
)

//此处用sync.once用来防止多次请求导致的内存浪费，因为register仅仅是用来调用Do方法的，不存在因为多个请求导致的registerService的内存访问出现的问题

func NewHandleRegister(user *Repository.User) (Controller.StatusCode, error) {
	registerOnce.Do(func() {
		registerService = &register{}
	})
	return registerService.Do(user)
}

//此处的两个操作不满足并发模型

func (service *register) Do(user *Repository.User) (Controller.StatusCode, error) {
	dao := Dao.NewUserDao()
	exist, err := dao.UsernameIsExist(user.Username)
	if exist {
		return UsernameExist, fmt.Errorf("用户名已经存在")
	}
	if err != nil {
		return ServerError, err
	}
	err = dao.KeepUserToDataSource(user)
	return SuccessRegister, err
}
