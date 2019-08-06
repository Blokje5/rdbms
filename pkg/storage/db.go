package storage

import (
	"os"
)

type DB struct {
	dbFile *os.File
}

// OpenDB initialises a DB against the file
func OpenDB(filename string) (*DB, error) {
	f, err := os.OpenFile(filename, os.O_RDWR | os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}

	db := &DB {
		dbFile: f,
	}

	return db, nil
}

