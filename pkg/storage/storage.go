package storage

type Storage interface {
	Get(file string) (string, error)
	Put(file string, content string) error
	Folder() string
}
