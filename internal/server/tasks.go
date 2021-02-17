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
	type request struct {
		Task string `json:"task"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var (
			res response
			req request
			err error
		)
		err = s.decode(w, r, &req)
		if err != nil {
			s.Logger.Printf("%s\n", err)
      s.respond(w, r, nil, http.StatusBadRequest)
		}

		// TODO validate input

		res.ID, err = s.TaskCreate(req.Task)
		if err != nil {
			s.Logger.Printf("db error: %s\n", err)
			http.Error(w, "db error", http.StatusInternalServerError)
		}
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
    s.respond(w, r, nil, http.StatusOK)
	}
}

func mapToXTask(m map[int]string) []Task {
	tasks := []Task{}
	for k, v := range m {
		tasks = append(tasks, Task{ID: k, Name: v})
	}
	return tasks
}
