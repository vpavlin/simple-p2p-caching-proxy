package store

type IStore interface {
	Get(key []byte) ([]byte, error)
	Put(key []byte, value []byte) error
	Has(key []byte) (bool, error)
}
