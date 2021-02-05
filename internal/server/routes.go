package server

import (
	"net/http"
	"os"
)

// Routes sets up all routes in the server.
func (s *Server) Routes() {
	r := s.Router

	assetsDir := "assets"
	_, err := os.Stat(assetsDir)
	if err != nil {
		s.Logger.Printf("%s\n", err)
  }
  // Serves files under http://localhost:1729/assets/<filename>
  r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir(assetsDir))))
  
  r.HandleFunc("/", s.HandleHome())
}
