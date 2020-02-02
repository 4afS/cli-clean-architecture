package infrastructure

import (
	bolt "go.etcd.io/bbolt"
	"log"
)

type DbHandler struct {
	Db *bolt.DB
}

func NewDbHandler(path string) *DbHandler {
	b, err := bolt.Open(path, 0666, nil)
	if err != nil {
		log.Fatal(err)
	}

	return &DbHandler{Db: b}
}
