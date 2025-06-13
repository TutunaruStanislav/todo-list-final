package api

import (
	"net/http"
	"strconv"

	"gop/pkg/db"
)

type TaskAddReponse struct {
	ID string `json:"id"`
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task *db.Task

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

	writeJson(w, TaskAddReponse{ID: strconv.FormatInt(id, 10)}, http.StatusOK)
}
