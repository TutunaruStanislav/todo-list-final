package api

import (
	"net/http"
	"time"

	"gop/pkg/db"
)

func taskDoneHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeError(w, err.Error(), http.StatusNotFound)
		return
	}

	if len(task.Repeat) == 0 {
		err = db.DeleteTask(id)
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
		err = db.UpdateTask(task)
		if err != nil {
			writeError(w, err.Error(), http.StatusNotFound)
			return
		}
	}

	writeJson(w, SuccessResponse, http.StatusOK)
}
