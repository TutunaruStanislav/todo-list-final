package api

import (
	"database/sql"
	"net/http"
	"time"

	"gop/pkg/db"
)

type TaskDoneHandler struct {
	db *sql.DB
}

func NewTaskDoneHandler(db *sql.DB) *TaskDoneHandler {
	return &TaskDoneHandler{
		db: db,
	}
}

func (h *TaskDoneHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	if len(task.Repeat) == 0 {
		err = db.DeleteTask(h.db, id)
		if err != nil {
			writeError(w, err.Error(), http.StatusNotFound)
			return
		}
	} else {
		task.Date, err = NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			writeError(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		err = db.UpdateTask(h.db, task)
		if err != nil {
			writeError(w, err.Error(), http.StatusNotFound)
			return
		}
	}

	writeJson(w, SuccessResponse, http.StatusOK)
}
