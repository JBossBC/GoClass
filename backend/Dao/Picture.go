package Dao

import (
	"goClass/backend/Repository"
	"io"
	"mime/multipart"
	"os"
	"sync"
)

var (
	pictureDao  *PictureDao
	pictureOnce sync.Once
)

type PictureDao struct {
}

func NewPictureDao() *PictureDao {
	pictureOnce.Do(func() {
		GetMysqlConnection().AutoMigrate(&Repository.Picture{})
		pictureDao = &PictureDao{}
	})
	return pictureDao
}
func (pictureDao *PictureDao) KeepPictureToMysql(picture *Repository.Picture) error {
	connection := GetMysqlConnection()
	return connection.Create(picture).Error
}
func (pictureDao *PictureDao) KeepPictureToLocal(fileHead *multipart.FileHeader, url string) (*os.File, error) {
	file, err := fileHead.Open()
	if err != nil {
		return nil, err
	}
	create, err := os.Create("picture\\" + url + ".jpg")
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(create, file)
	if err != nil {
		return nil, err
	}
	return create, nil
}
