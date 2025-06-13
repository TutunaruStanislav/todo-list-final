package api

import (
	"net/http"
	"strconv"

	"gop/pkg/db"
)

type Success struct {
	ID string `json:"id"`
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task *db.Task
	var success Success

	task, err := validateRequest(r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := db.AddTask(task)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	success.ID = strconv.FormatInt(id, 10)

	writeJson(w, success, http.StatusOK)
}
