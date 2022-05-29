package Repository

import gorm2 "gorm.io/gorm"

type Picture struct {
	gorm2.Model
	Url string
}
