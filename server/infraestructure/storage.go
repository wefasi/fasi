package infraestructure

import (
	"github.com/wefasi/fasi/server/infraestructure/storage"
)

var s3 storage.Storage

func InitS3() {
	_s3 := storage.NewS3Storage()
	s3 = &_s3
}

func GetS3() storage.Storage {
	return s3
}
