package Service

import (
	"goClass/backend/Dao"
	"goClass/backend/Repository"
	"sync"
)

type log struct {
}

var (
	logService *log
	logOnce    sync.Once
)

func NewLogHandle(user *Repository.User) (bool, error) {
	logOnce.Do(func() {
		logService = &log{}
	})
	return logService.Do(user)
}
func (logService *log) Do(user *Repository.User) (bool, error) {
	exist, err := Dao.NewUserDao().UserIsExist(user)
	if err != nil {
		return false, err
	}
	if exist {
		return true, nil
	}
	return false, nil
}
