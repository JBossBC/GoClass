package Controller

import (
	"goClass/backend/Service"
	"mime/multipart"
)

func UploadPicture(picture *multipart.FileHeader) error {

	err := Service.NewPicture().UploadPicture(picture)
	return err
}
