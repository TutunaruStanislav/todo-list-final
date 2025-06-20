package api

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"gop/pkg/db"
)

type TaskAddResponse struct {
	ID string `json:"id"`
}

type AddTaskHandler struct {
	db *sql.DB
}

func NewAddTaskHandler(db *sql.DB) *AddTaskHandler {
	return &AddTaskHandler{
		db: db,
	}
}

// AddTaskHandler is a POST request handler /api/addtask.

// It receives the request data, validates it, adds the record to the DB
// and returns the record id if successful, otherwise an error.
func (h *AddTaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var task *db.Task

	task, err := validateRequest(r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := db.AddTask(h.db, task)
	if err != nil {
		log.Println("AddTask:", err)
		writeError(w, InternalServerErrorMessage, http.StatusInternalServerError)
		return
	}

	writeJson(w, TaskAddResponse{ID: strconv.FormatInt(id, 10)}, http.StatusOK)
}
