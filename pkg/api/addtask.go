package api

import (
	"database/sql"
	"net/http"
	"strconv"

	"gop/pkg/db"
)

type TaskAddReponse struct {
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

func (h *AddTaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var task *db.Task

	task, err := validateRequest(r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := db.AddTask(h.db, task)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJson(w, TaskAddReponse{ID: strconv.FormatInt(id, 10)}, http.StatusOK)
}
