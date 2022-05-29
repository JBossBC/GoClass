package Service

import (
	"goClass/backend/Dao"
	"goClass/backend/Repository"
	"goClass/backend/util"
	gorm2 "gorm.io/gorm"
	"mime/multipart"
	"sync"
)

var (
	pictureServer *PictureServer
	pictureOnce   sync.Once
)

type PictureServer struct {
}

func NewPicture() *PictureServer {
	pictureOnce.Do(func() {
		pictureServer = &PictureServer{}
	})
	return pictureServer
}
func (pictureServer *PictureServer) UploadPicture(pictureFileHeader *multipart.FileHeader) error {
	url := util.MD5EnCrypto(pictureFileHeader.Filename)

	local, err := Dao.NewPictureDao().KeepPictureToLocal(pictureFileHeader, url)
	if err != nil {
		return err
	}
	//这里有一点漏洞，适用性不强，应该使用当前工作路径
	picture := &Repository.Picture{
		Model: gorm2.Model{},
		Url:   local.Name(),
	}
	err = Dao.NewPictureDao().KeepPictureToMysql(picture)
	return err
}
