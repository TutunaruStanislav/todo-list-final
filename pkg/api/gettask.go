package api

import (
	"database/sql"
	"net/http"

	"gop/pkg/db"
)

type GetTaskHandler struct {
	db *sql.DB
}

func NewGetTaskHandler(db *sql.DB) *GetTaskHandler {
	return &GetTaskHandler{
		db: db,
	}
}

func (h *GetTaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := db.GetTask(h.db, id)
	if err != nil {
		writeError(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJson(w, task, http.StatusOK)
}
