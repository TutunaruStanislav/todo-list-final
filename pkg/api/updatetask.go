package api

import (
	"database/sql"
	"log"
	"net/http"

	"gop/pkg/db"
)

type UpdateTaskHandler struct {
	db *sql.DB
}

func NewUpdateTaskHandler(db *sql.DB) *UpdateTaskHandler {
	return &UpdateTaskHandler{
		db: db,
	}
}

// UpdateTaskHandler is a handler for the /api/task PUT request.
//
// It receives parameters, validates, updates the DB record with task information
// and returns {} if successful, otherwise an error.
func (h *UpdateTaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	task, err := validateRequest(r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.UpdateTask(h.db, task)
	if err != nil {
		if err == sql.ErrNoRows {
			writeError(w, "incorrect id for updating task", http.StatusNotFound)
			return
		} else {
			log.Println("UpdateTask:", err)
			writeError(w, InternalServerErrorMessage, http.StatusInternalServerError)
			return
		}
	}
	writeJson(w, SuccessResponse, http.StatusOK)
}
