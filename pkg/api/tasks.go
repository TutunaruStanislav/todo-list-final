package api

import (
	"net/http"

	"gop/pkg/db"
)

const maxTasks = 10

type TasksResponse struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.GetTasks(maxTasks)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
	}
	writeJson(w, TasksResponse{Tasks: tasks}, http.StatusOK)
}
