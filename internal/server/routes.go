package server

import (
	"net/http"
	"os"
)

// Routes sets up all routes in the server.
func (s *Server) Routes() {
  r := s.Router
  var err error

  // Setup Fileserver
  // fsDir is a simple struct that is only used to make fileserver setup easier.
	type fsDir struct {
		prefix string
		dir    string
  }
  // NOTE Append new fileservers here
  var fsDirs []fsDir
  fsDirs = append(fsDirs, fsDir{prefix: "/assets/", dir: "assets"})
  fsDirs = append(fsDirs, fsDir{prefix: "/app/", dir: "web/app"})
  fsDirs = append(fsDirs, fsDir{prefix: "/static/", dir: "web/static"})
  for _, v := range fsDirs {
    // Check if directory exists
    _, err = os.Stat(v.dir)
    if err != nil {
      s.Logger.Printf("%s\n", err)
    }
    // Add to router
    r.PathPrefix(v.prefix).Handler(http.StripPrefix(v.prefix, http.FileServer(http.Dir(v.dir))))
  }

	r.HandleFunc("/", s.HandleHome())
}