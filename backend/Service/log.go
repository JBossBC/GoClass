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

func NewLogHandle(user *Repository.User) bool {
	logOnce.Do(func() {
		logService = &log{}
	})
	return logService.Do(user)
}
func (logService *log) Do(user *Repository.User) bool {
	exist := Dao.NewUserDao().UserIsExist(user)
	if exist {
		return true
	}
	return false
}
