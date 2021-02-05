package server

import (
	"html/template"
	"log"
	"net/http"
	"os"
	// bolt "go.etcd.io/bbolt"
)

// Run will start the homepage server.
func Run() error {
	var err error

	s := NewServer()
	http.ListenAndServe(":1729", s)

	return err
}

// Server holds the handler, database, and methods used by homepage.
type Server struct {
	Logger *log.Logger
	// DB      *bolt.DB
}

// NewServer is a constructor function for Server.
func NewServer() *Server {
	s := Server{
		Logger: log.New(os.Stdout, "homepage: ", log.Lshortfile),
	}

	// TODO implement DB setup

	return &s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.New("").ParseFiles("web/main.tpl")
	if err != nil {
    s.Logger.Fatalf("%s\n", err)
  }
  
  err = tpl.ExecuteTemplate(w, "main", nil)
  if err != nil {
    s.Logger.Fatalf("%s\n", err)
  }
}
