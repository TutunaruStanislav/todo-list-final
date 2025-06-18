package api

import (
	"database/sql"
	"net/http"

	"gop/pkg/db"
)

type DeleteTaskHandler struct {
	db *sql.DB
}

func NewDeleteTaskHandler(db *sql.DB) *DeleteTaskHandler {
	return &DeleteTaskHandler{
		db: db,
	}
}

// DeleteTaskHandler is a handler for the DELETE request /api/task?id=<id>, where <id> is the task ID.

// It gets the task id from GET parameters, validates and permanently deletes the task from DB.

// It returns {} if successful, otherwise an error.
func (h *DeleteTaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.DeleteTask(h.db, id)
	if err != nil {
		writeError(w, err.Error(), http.StatusNotFound)
		return
	}

	writeJson(w, SuccessResponse, http.StatusOK)
}
