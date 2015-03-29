package store

import (
	"os"
)

/*
* Generic Backend Store interface and associated interface methods that this
* interface expose to handle the task in our hand.
 */
type Store interface {
	// Initilizes the store after supplying the configurations.
	Init()

	// Store provides an atomic interface to make the changes atomically.
	AcquireLock() error
	ReleaseLock() error

	// CURD operations.
	AddObject(path string, obj []byte) error
	AddFile(path string, file *os.File, contentType string) error
	ReadAsObject(path string) ([]byte, error)
	ReadAsFile(path string) error
}

/*
* Error interface for the Store.
 */
type StoreError struct {
	StoreId string
}

func (store *StoreError) Error() string {
	return store.StoreId
}
