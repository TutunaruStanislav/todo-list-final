package api

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:
		addTaskHandler(w, r)
	}
}

func Init() {
	http.HandleFunc("/api/nextdate", nextDayHandler)
	http.HandleFunc("/api/task", taskHandler)
	http.HandleFunc("/api/tasks", tasksHandler)
}

func writeJson(w http.ResponseWriter, data any, statusCode int) {
	resp, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	w.Write(resp)
}

func writeError(w http.ResponseWriter, err string, statusCode int) {
	var error Error
	error.Error = err
	writeJson(w, error, statusCode)
}
