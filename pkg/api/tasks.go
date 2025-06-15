package api

import (
	"database/sql"
	"net/http"
	"time"

	"gop/pkg/db"
)

const maxTasks = 10

type TasksResponse struct {
	Tasks []*db.Task `json:"tasks"`
}

type TasksHandler struct {
	db *sql.DB
}

func NewTasksHandler(db *sql.DB) *TasksHandler {
	return &TasksHandler{
		db: db,
	}
}

func (h *TasksHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	var date string
	if len(search) > 0 {
		parsedDate, err := time.Parse(InputDateFormat, search)
		if err == nil {
			date = parsedDate.Format(DateFormat)
			search = ""
		}
	}

	tasks, err := db.GetTasks(h.db, maxTasks, search, date)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
	}
	writeJson(w, TasksResponse{Tasks: tasks}, http.StatusOK)
}

// func tasksHandler(w http.ResponseWriter, r *http.Request) {
// 	search := r.URL.Query().Get("search")
// 	var date string
// 	if len(search) > 0 {
// 		parsedDate, err := time.Parse(InputDateFormat, search)
// 		if err == nil {
// 			date = parsedDate.Format(DateFormat)
// 			search = ""
// 		}
// 	}

// 	tasks, err := db.GetTasks(maxTasks, search, date)
// 	if err != nil {
// 		writeError(w, err.Error(), http.StatusInternalServerError)
// 	}
// 	writeJson(w, TasksResponse{Tasks: tasks}, http.StatusOK)
// }
