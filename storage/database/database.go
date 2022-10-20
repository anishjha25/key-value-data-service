package database

type DB interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

func NewDB() DB {
	return newMemoryDB()
}
