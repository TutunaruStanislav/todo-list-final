package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

var SuccessResponse struct{}

type TaskHandler struct {
	db *sql.DB
}

func NewTaskHandler(db *sql.DB) *TaskHandler {
	return &TaskHandler{
		db: db,
	}
}

func (h *TaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:
		NewAddTaskHandler(h.db).ServeHTTP(w, r)
	case http.MethodGet:
		NewGetTaskHandler(h.db).ServeHTTP(w, r)
	case http.MethodPut:
		NewUpdateTaskHandler(h.db).ServeHTTP(w, r)
	case http.MethodDelete:
		NewDeleteTaskHandler(h.db).ServeHTTP(w, r)
	}
}

func Init(db *sql.DB) {
	http.HandleFunc("/api/nextdate", nextDayHandler)
	http.HandleFunc("/api/signin", signInHandler)
	http.HandleFunc("/api/task", auth(NewTaskHandler(db).ServeHTTP))
	http.HandleFunc("/api/tasks", auth(NewTasksHandler(db).ServeHTTP))
	http.HandleFunc("/api/task/done", auth(NewTaskDoneHandler(db).ServeHTTP))
}

func writeJson(w http.ResponseWriter, data any, statusCode int) {
	resp, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	w.Write(resp)
}

func writeError(w http.ResponseWriter, err string, statusCode int) {
	writeJson(w, Error{Error: err}, statusCode)
}
