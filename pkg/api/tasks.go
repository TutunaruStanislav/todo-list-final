package api

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"gop/pkg/db"
)

const maxTasks = 10 // tasks limit

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

// TasksHandler is a handler for the GET request /api/tasks?search=<search>,
// where <search> is the string to search for.
//
// It finds up to 10 tasks filtered by <search> (if passed) and sorted
// by date in ascending order and returns them, otherwise an error.
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
		log.Println("GetTasks:", err)
		writeError(w, InternalServerErrorMessage, http.StatusInternalServerError)
		return
	}
	writeJson(w, TasksResponse{Tasks: tasks}, http.StatusOK)
}
