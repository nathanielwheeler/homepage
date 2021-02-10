package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"
	bolt "go.etcd.io/bbolt"
)

// Run will start the homepage server.
func Run() error {
	var err error

	s := NewServer()
	port := ":1729"
	fmt.Printf("Listening on %s...\n", port)
	http.ListenAndServe(port, s)

	return err
}

// Server holds the handler, database, and methods used by homepage.
type Server struct {
	Logger *log.Logger
	Router *mux.Router
	DB     *bolt.DB
}

// NewServer is a constructor function for Server.
func NewServer() *Server {
	s := Server{
		Logger: log.New(os.Stdout, "homepage: ", log.Lshortfile),
		Router: mux.NewRouter(),
	}

	s.Routes()
	err := s.DBInit("tools/bbolt.db")
	if err != nil {
		s.Logger.Printf("%s\n", err)
	}

	return &s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

// HandleHome will serve the home page
func (s *Server) HandleHome() http.HandlerFunc {
	var (
		init sync.Once
		tpl  *template.Template
		err  error
	)
	return func(w http.ResponseWriter, r *http.Request) {
		init.Do(func() {
			tpl, err = template.New("").ParseGlob("web/template/*.tpl")
		})
		if err != nil {
			s.Logger.Printf("%s\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tpl.ExecuteTemplate(w, "main", nil)
		if err != nil {
			s.Logger.Fatalf("%s\n", err)
		}
	}

}

// HandleError will return an error message.
func (s *Server) HandleError(msg string, code int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, msg, code)
	}
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, data interface{}, code int) {
	w.WriteHeader(code)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			s.Logger.Printf("%s\n", err)
			http.Error(w, "encoding error", http.StatusInternalServerError)
		}
	}
}

func (s *Server) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
