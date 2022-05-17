package Controller

import (
	"fmt"
	"goClass/backend/Repository"
	"goClass/backend/Service"
	"goClass/backend/util"
)

//检查第一次输入的密码和第二次密码的方式最好放在前端执行
const (
	InputPasswordError Service.StatusCode = 401
)

func Register(username string, password string, rePassword string) Service.StatusCode {
	//controller层处理username和password以及rePassword

	//先对password和rePassword进行验证处理
	if password != rePassword {
		fmt.Println("the passwords you entered twice are inconsistent")
		return InputPasswordError
	}
	//对User进行封装
	user := Repository.User{
		Username: username,
		Password: util.MD5EnCrypto(password),
	}
	//交给Service进行处理
	register, err := Service.NewHandleRegister(&user)
	if err != nil {
		fmt.Println(err)
	}
	return register
}

//func Log(username string, password string)(StatusCode, error){
//
//
//
//	return
//}
