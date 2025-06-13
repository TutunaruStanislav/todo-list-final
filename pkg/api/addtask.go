package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"gop/pkg/db"
)

type Success struct {
	ID string `json:"id"`
}

func checkDate(task *db.Task) error {
	now := time.Now()
	if len(task.Date) == 0 {
		task.Date = now.Format(DateFormat)
	}

	t, err := time.Parse(DateFormat, task.Date)
	if err != nil {
		return err
	}

	if afterNow(now, t) {
		if len(task.Repeat) > 0 {
			nextDate, err := NextDate(now, task.Date, task.Repeat)
			if err != nil {
				return err
			}
			task.Date = nextDate
		} else {
			task.Date = now.Format(DateFormat)
		}
	}

	return nil
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	var success Success
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(task.Title) == 0 {
		writeError(w, "title cannot be blank", http.StatusBadRequest)
		return
	}

	err = checkDate(&task)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	success.ID = strconv.FormatInt(id, 10)

	writeJson(w, success, http.StatusOK)
}
