package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"gop/pkg/db"
)

// checkDate - validates and updates the date in the task if successful, otherwise returns an error.
func checkDate(task *db.Task) error {
	now := time.Now()
	if len(task.Date) == 0 {
		task.Date = now.Format(DateFormat)
	}

	t, err := time.Parse(DateFormat, task.Date)
	if err != nil {
		return errors.New("incorrect format of the date")
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

// validateRequest - parses and validates the request, if successful it returns the filled task model, otherwise it returns an error.
func validateRequest(r *http.Request) (*db.Task, error) {
	var buf bytes.Buffer
	task := &db.Task{}

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
		return nil, err
	}

	if len(task.Title) == 0 {
		return nil, errors.New("title cannot be blank")
	}

	err = checkDate(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// parseId - parses and validates id from string parameter, in case of success returns id int64, otherwise error.
func parseId(r *http.Request) (int64, error) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		return 0, errors.New("id cannot be blank")
	}

	return id, nil
}

// validatePassword - parses and validates the parameters from the request, if successful returns the filled user model, otherwise error.
func validatePassword(r *http.Request) (*db.User, error) {
	var buf bytes.Buffer
	var input db.UserInput

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(buf.Bytes(), &input); err != nil {
		return nil, err
	}

	user := input.ToUser()
	if len(user.Password) == 0 {
		return nil, errors.New("password cannot be blank")
	}

	return user, nil
}
