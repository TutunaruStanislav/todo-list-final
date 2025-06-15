package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Error struct {
	Error string `json:"error"`
}

var SuccessResponse struct{}

func Init(db *sql.DB, router *chi.Mux) {
	router.Get("/api/nextdate", nextDayHandler)
	router.Post("/api/signin", signInHandler)
	router.Post("/api/task/done", auth(NewTaskDoneHandler(db).ServeHTTP))
	router.Get("/api/tasks", auth(NewTasksHandler(db).ServeHTTP))
	router.Get("/api/task", auth(NewGetTaskHandler(db).ServeHTTP))
	router.Post("/api/task", auth(NewAddTaskHandler(db).ServeHTTP))
	router.Put("/api/task", auth(NewUpdateTaskHandler(db).ServeHTTP))
	router.Delete("/api/task", auth(NewDeleteTaskHandler(db).ServeHTTP))
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
