package server

import (
	"encoding/binary"
	"time"

	bolt "go.etcd.io/bbolt"
)

// DBInit initializes a bbolt database.
func (s *Server) DBInit(dbPath string) error {
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	buckets := []string{
		"tasks",
	}

	for _, v := range buckets {
		// Create Buckets if they don't exist
		err = db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(v))
			return err
		})
		if err != nil {
			return err
		}
	}

	s.DB = db

	return nil
}

// TaskGetAll will return all tasks as a map.
func (s *Server) TaskGetAll() (map[int]string, error) {
	tasks := make(map[int]string)
	err := s.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks[btoi(k)] = string(v)
		}
		return nil
  })
  if err != nil {
    return nil, err
  }
	return tasks, nil
}

// TaskCreate will make a new task and return its key, returning -1 and an error if there is a transaction problem.  The actual key within the database is stored as a big-endian uint64.
func (s *Server) TaskCreate(task string) (int, error) {
	var id int
	err := s.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		uid64, _ := b.NextSequence()
		id = int(uid64)
		return b.Put(utob(uid64), []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

// TaskDelete will remove a task from the database.
func (s *Server) TaskDelete(key int) error {
  return s.DB.Update(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte("tasks"))
    return b.Delete(utob(uint64(key)))
  })
}

func utob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
