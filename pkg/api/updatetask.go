package api

import (
	"gop/pkg/db"
	"net/http"
)

var successUpdate struct{}

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	task, err := validateRequest(r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.UpdateTask(task)
	if err != nil {
		writeError(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJson(w, successUpdate, http.StatusOK)
}
