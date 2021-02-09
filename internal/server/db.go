package server

import (
	"time"

	bolt "go.etcd.io/bbolt"
)

// DBInit initializes a bbolt database.
func (s *Server) DBInit(dbPath string) error {
  db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
  if err != nil {
    return err
  }
  err = db.Update(func(tx *bolt.Tx) error {
    _, err := tx.CreateBucketIfNotExists([]byte("tasks"))
    return err
  })
  if err != nil {
    return err
  }
  s.DB = db

  return nil
}
