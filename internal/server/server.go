package server

import (
	"log"
	"net/http"
	"os"
	// bolt "go.etcd.io/bbolt"
)

// Run will start the homepage server.
func Run() error {
	var err error

	s := NewServer()
	http.ListenAndServe(":1729", *s.Handler)

	return err
}

// Server holds the handler, database, and methods used by homepage.
type Server struct {
	Handler *http.Handler
	Logger  *log.Logger
	// DB      *bolt.DB
}

// NewServer is a constructor function for Server.
func NewServer() *Server {
	s := Server{
		Logger: log.New(os.Stdout, "homepage: ", log.Lshortfile),
  }
  
  // TODO implement handler

	// TODO implement DB setup

	return &s
}
