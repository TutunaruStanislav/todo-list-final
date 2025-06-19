package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const InternalServerErrorMessage = "internal server error"

type Error struct {
	Error string `json:"error"`
}

var SuccessResponse struct{}

// Init registers handlers for all api requests.
func Init(db *sql.DB, router *chi.Mux) {
	router.Get("/api/nextdate", NextDayHandler)
	router.Post("/api/signin", SignInHandler)
	router.Post("/api/task/done", auth(NewTaskDoneHandler(db).ServeHTTP))
	router.Get("/api/tasks", auth(NewTasksHandler(db).ServeHTTP))
	router.Get("/api/task", auth(NewGetTaskHandler(db).ServeHTTP))
	router.Post("/api/task", auth(NewAddTaskHandler(db).ServeHTTP))
	router.Put("/api/task", auth(NewUpdateTaskHandler(db).ServeHTTP))
	router.Delete("/api/task", auth(NewDeleteTaskHandler(db).ServeHTTP))
}

// writeJson - serializes and gives data in JSON format to the response.
func writeJson(w http.ResponseWriter, data any, statusCode int) {
	resp, err := json.Marshal(data)
	if err != nil {
		log.Println("serializeJson:", err)
		http.Error(w, InternalServerErrorMessage, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	_, err = w.Write(resp)
	if err != nil {
		log.Println("writeJson:", err)
	}
}

// writeError is a wrapper over the error.
func writeError(w http.ResponseWriter, err string, statusCode int) {
	writeJson(w, Error{Error: err}, statusCode)
}
