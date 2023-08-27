package infraestructure

import (
	"time"

	"github.com/wefasi/fasi/server/infraestructure/storage"
)

var cache storage.Storage

func InitCache() {
	c := storage.NewLocalDB(time.Hour)
	cache = &c
}

func GetCache() storage.Storage {
	return cache
}
