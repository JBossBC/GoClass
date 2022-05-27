package Repository

import (
	gorm2 "gorm.io/gorm"
)

type Article struct {
	gorm2.Model
	//加快搜索
	UserName string `gorm:"index"`
	Header   string
	Context  string
}
