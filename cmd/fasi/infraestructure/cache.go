package infraestructure

import (
	"time"

	pkg "github.com/wefasi/fasi/pkg/storage"
)

var cache pkg.Storage

func InitCache() {
	c := pkg.NewLocalDB(time.Minute)
	cache = &c
}

func GetCache() pkg.Storage {
	return cache
}
