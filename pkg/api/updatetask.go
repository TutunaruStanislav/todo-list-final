package api

import (
	"database/sql"
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
		writeError(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJson(w, SuccessResponse, http.StatusOK)
}
