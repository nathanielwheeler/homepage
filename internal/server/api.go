package server

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// HandleTasks will handle API calls related to tasks.
func (s *Server) HandleTasks(cmd string) http.HandlerFunc {
	switch cmd {
	case "create":
    return func(w http.ResponseWriter, r *http.Request) {
      // TODO implement
    }
	case "delete":
		return func(w http.ResponseWriter, r *http.Request) {
			id, err := strconv.Atoi(mux.Vars(r)["id"])
			if err != nil {
				s.Logger.Printf("%s\n", err)
				http.Error(w, "invalid ID", http.StatusBadRequest)
			}
			err = s.TaskDelete(id)
      if err != nil {
        s.Logger.Printf("%s\n", err)
        http.Error(w, "db error", http.StatusInternalServerError)
      }
		}
	}
	// This should never be reached.
	s.Logger.Println("HandleTasks called with invalid string")
	return s.HandleError("get thee back, ghost", http.StatusForbidden)
}
