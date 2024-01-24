package store

type Store interface {
	Exist(string) (bool, error)
	Set(string) error
	Delete(string) error
}
