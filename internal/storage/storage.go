package storage

type Storage interface {
	Find(items []uint64) ([]uint64, error)
	Add(items []uint64) error
	Delete(items []uint64) error
}
