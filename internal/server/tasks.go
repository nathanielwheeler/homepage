package server

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Task holds task metadata
type Task struct {
	ID   int
	Name string
}

// HandleTaskCreate will handle requests to create tasks.
func (s *Server) HandleTaskCreate() http.HandlerFunc {
	type response struct {
		ID int `json:"id"`
	}
  // FIXME receiving empty request from client
	type request struct {
		Task string `json:"task"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			res response
			req request
			err error
		)
		s.decode(w, r, req)
		s.Logger.Println(req)
		res.ID, err = s.TaskCreate(req.Task)
		if err != nil {
			s.Logger.Printf("db error: %s\n", err)
			http.Error(w, "db error", http.StatusInternalServerError)
		}
		s.Logger.Println("task id:", res.ID)
		s.respond(w, r, res, http.StatusOK)
	}
}

// HandleTaskDelete will handle requests to delete tasks.
func (s *Server) HandleTaskDelete() http.HandlerFunc {

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

func mapToXTask(m map[int]string) []Task {
	tasks := []Task{}
	for k, v := range m {
		tasks = append(tasks, Task{ID: k, Name: v})
	}
	return tasks
}
