package api

import (
	"net/http"
	"time"

	"gop/pkg/db"
)

const maxTasks = 10

type TasksResponse struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	var date string
	if len(search) > 0 {
		parsedDate, err := time.Parse(InputDateFormat, search)
		if err == nil {
			date = parsedDate.Format(DateFormat)
			search = ""
		}
	}

	tasks, err := db.GetTasks(maxTasks, search, date)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
	}
	writeJson(w, TasksResponse{Tasks: tasks}, http.StatusOK)
}
