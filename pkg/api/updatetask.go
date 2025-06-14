package api

import (
	"net/http"

	"gop/pkg/db"
)

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
	writeJson(w, SuccessResponse, http.StatusOK)
}
