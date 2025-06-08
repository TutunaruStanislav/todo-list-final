package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const DateFormat = "20060102"
const maxDaysInterval = 400

func afterNow(date, now time.Time) bool {
	return date.After(now)
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	date, err := time.Parse(DateFormat, dstart)
	if err != nil {
		return "", err
	}

	if repeat == "" {
		return "", errors.New("incorrect repeat rule provided")
	}

	chunks := strings.Split(repeat, " ")
	if (chunks[0] == "y" && len(chunks) > 1) || (chunks[0] == "d" && len(chunks) == 1) {
		return "", errors.New("incorrect repeat rule provided")
	}

	if chunks[0] == "y" || chunks[0] == "d" {
		var dayInterval int
		if chunks[0] == "d" {
			dayInterval, err = strconv.Atoi(chunks[1])
			if err != nil {
				return "", err
			}
			if dayInterval > maxDaysInterval {
				return "", errors.New("max days interval was overlimited")
			}
		}

		for {
			if chunks[0] == "y" {
				date = date.AddDate(1, 0, 0)
			} else {
				date = date.AddDate(0, 0, dayInterval)
			}
			if afterNow(date, now) {
				break
			}
		}

		return date.Format(DateFormat), nil
	}

	return "", errors.New("incorrect repeat rule provided")
}

func nextDayHandler(w http.ResponseWriter, r *http.Request) {
	now := r.FormValue("now")
	currentTime, err := time.Parse(DateFormat, now)
	if err != nil {
		currentTime = time.Now()
	}

	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	nextDate, err := NextDate(currentTime, date, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(nextDate))
}
