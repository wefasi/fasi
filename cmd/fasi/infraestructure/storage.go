package infraestructure

import (
	"github.com/wefasi/fasi/pkg/storage"
)

var s3 *storage.S3Storage

func InitS3() {
	_s3 := storage.NewS3Storage("")
	s3 = &_s3
}

func GetS3() *storage.S3Storage {
	return s3
}
