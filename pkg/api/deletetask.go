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
