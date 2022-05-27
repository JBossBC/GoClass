package Repository

import (
	gorm2 "gorm.io/gorm"
)

type User struct {
	gorm2.Model
	Username string
	Password string
}
