package infraestructure

import (
	pkg "github.com/wefasi/fasi/pkg/storage"
)

var s3 pkg.Storage

func InitS3() {
	_s3 := pkg.NewS3Storage()
	s3 = &_s3
}

func GetS3() pkg.Storage {
	return s3
}
