package Service

import (
	"goClass/backend/Dao"
	"goClass/backend/util"
	"mime/multipart"
	"sync"
)

var (
	pictureServer *PictureServer
	pictureOnce   *sync.Once
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

	Dao.NewPictureDao().KeepPictureToLocal(pictureFileHeader, url)
	return nil
}
