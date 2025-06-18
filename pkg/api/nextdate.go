package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const DateFormat = "20060102"        // main date form
const InputDateFormat = "02.01.2006" // date format for parse input search parameter
const maxDaysInterval = 400          // maximum interval of task transfer in days

var daysArray [32]bool
var monthsArray [13]bool

// datesAreSame is a function that compares the current date with the last or penultimate day of the month, depending on the rules.
func datesAreSame(date time.Time, firstOfMonth time.Time, last bool, prevLast bool) bool {
	if last && prevLast {
		lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
		prevLastOfMonth := firstOfMonth.AddDate(0, 1, -2)
		if date.Day() == lastOfMonth.Day() || date.Day() == prevLastOfMonth.Day() {
			return true
		}
	} else if last {
		lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
		if date.Day() == lastOfMonth.Day() {
			return true
		}
	} else if prevLast {
		prevLastOfMonth := firstOfMonth.AddDate(0, 1, -2)
		if date.Day() == prevLastOfMonth.Day() {
			return true
		}
	}

	return false
}

// weekdaysAreSame is a function that compares weekday numbers.
func weekdaysAreSame(date time.Time, parsedWeekday int) bool {
	currentWeekday := int(date.Weekday())
	if currentWeekday == 0 {
		currentWeekday = 7
	}
	if currentWeekday == parsedWeekday {
		return true
	}

	return false
}

// afterNow is a function that compares 2 dates and returns true if the first is greater than the second.
func afterNow(date time.Time, now time.Time) bool {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC).After(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC))
}

// fillDaysArray - function for filling the array of days.
func fillDaysArray(day int) int {
	switch day {

	case -1:
		for index, _ := range monthsArray {
			monthsArray[index] = true
		}
	case -2:
		for index, _ := range monthsArray {
			monthsArray[index] = true
		}
	default:
		daysArray[day] = true
	}

	return day
}

// parseAndCheckWeekDay - A function that converts a string to a numeric number value for the week, validates and returns it if successful, otherwise an error.
func parseAndCheckWeekDay(day string) (int, error) {
	weekday, err := strconv.Atoi(day)
	if err != nil {
		return 0, errors.New("incorrect repeat rule provided")
	}
	if weekday < 1 || weekday > 7 {
		return 0, errors.New("incorrect repeat rule provided")
	}

	return weekday, nil
}

// parseAndCheckDay - a function that converts a string to a numeric value of the day of the month, validates and returns it if successful, otherwise an error.
func parseAndCheckDay(day string) (int, error) {
	currentDay, err := strconv.Atoi(day)
	if err != nil {
		return 0, errors.New("incorrect repeat rule provided")
	}
	if currentDay < -2 || currentDay == 0 || currentDay > 31 {
		return 0, errors.New("incorrect repeat rule provided")
	}

	return currentDay, nil
}

// parseAndCheckMonth - a function that converts a string to a numeric month value, validates and returns it if successful, otherwise an error.
func parseAndCheckMonth(month string) (int, error) {
	currentMonth, err := strconv.Atoi(month)
	if err != nil {
		return 0, errors.New("incorrect repeat rule provided")
	}
	if currentMonth == 0 || currentMonth > 12 {
		return 0, errors.New("incorrect repeat rule provided")
	}

	return currentMonth, nil
}

// parseDaysAndMonth - a function that parses rules specifying m (month), which returns an error if it fails.
func parseDaysAndMonth(days string, months string) (bool, bool, error) {
	var last bool
	var prevLast bool
	daysArray = [32]bool{}
	monthsArray = [13]bool{}

	daysChunks := strings.Split(days, ",")
	if len(daysChunks) == 1 {
		day, err := parseAndCheckDay(daysChunks[0])
		if err != nil {
			return false, false, err
		}

		res := fillDaysArray(day)
		if res == -1 {
			last = true
		}
		if res == -2 {
			prevLast = true
		}
	} else {
		for _, daysChunk := range daysChunks {
			day, err := parseAndCheckDay(daysChunk)
			if err != nil {
				return false, false, err
			}

			res := fillDaysArray(day)
			if res == -1 {
				last = true
			}
			if res == -2 {
				prevLast = true
			}
		}
	}

	if len(months) > 0 {
		monthsArray = [13]bool{}
		monthsChunks := strings.Split(months, ",")
		if len(monthsChunks) == 1 {
			month, err := parseAndCheckMonth(monthsChunks[0])
			if err != nil {
				return false, false, err
			}

			monthsArray[month] = true
		} else {
			for _, monthsChunk := range monthsChunks {
				month, err := parseAndCheckMonth(monthsChunk)
				if err != nil {
					return false, false, err
				}

				monthsArray[month] = true
			}
		}
	}

	return last, prevLast, nil
}

// NextDate is a function that calculates and returns a date string for the next task run if successful, otherwise an error.
//
// Parameters:
//   - now time.Time is the current date
//   - dstart string - task start date string
//   - repeat string - string with the rules of task repetition
func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	date, err := time.Parse(DateFormat, dstart)
	if err != nil {
		return "", err
	}

	if len(repeat) == 0 {
		return "", errors.New("incorrect repeat rule provided")
	}

	chunks := strings.Split(repeat, " ")
	switch chunks[0] {

	case "y", "d":
		if (chunks[0] == "y" && len(chunks) > 1) || (chunks[0] == "d" && len(chunks) == 1) {
			return "", errors.New("incorrect repeat rule provided")
		}

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

	case "m":
		switch len(chunks) {

		case 2:
			last, prevLast, err := parseDaysAndMonth(chunks[1], "")
			if err != nil {
				return "", err
			}
			for {
				date = date.AddDate(0, 0, 1)
				if afterNow(date, now) {
					if monthsArray[int(date.Month())] {
						firstOfMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
						if datesAreSame(date, firstOfMonth, last, prevLast) {
							break
						}
					}
					if daysArray[int(date.Day())] {
						break
					}
				}
			}
		case 3:
			last, prevLast, err := parseDaysAndMonth(chunks[1], chunks[2])
			if err != nil {
				return "", err
			}

			for {
				date = date.AddDate(0, 0, 1)
				if afterNow(date, now) {
					if monthsArray[int(date.Month())] {
						firstOfMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
						if daysArray[int(date.Day())] {
							break
						} else {
							if datesAreSame(date, firstOfMonth, last, prevLast) {
								break
							}
						}
					}
				}
			}

		default:
			return "", errors.New("incorrect repeat rule provided")
		}

	case "w":
		if len(chunks) < 2 {
			return "", errors.New("incorrect repeat rule provided")
		}

		weekdays := strings.Split(chunks[1], ",")
		if len(weekdays) == 1 {
			weekday, err := parseAndCheckWeekDay(weekdays[0])
			if err != nil {
				return "", err
			}

			for {
				date = date.AddDate(0, 0, 1)
				if afterNow(date, now) {
					if weekdaysAreSame(date, weekday) {
						break
					}
				}
			}
		} else {
			parsedWeekdays := []int{}
			for _, weekday := range weekdays {
				day, err := parseAndCheckWeekDay(weekday)
				if err != nil {
					return "", err
				}
				parsedWeekdays = append(parsedWeekdays, day)
			}

		outerLoop:
			for {
				date = date.AddDate(0, 0, 1)
				if afterNow(date, now) {
					for _, weekday := range parsedWeekdays {
						if weekdaysAreSame(date, weekday) {
							break outerLoop
						}
					}
				}
			}
		}
	default:
		return "", errors.New("incorrect repeat rule provided")
	}

	return date.Format(DateFormat), nil
}

// NextDayHandler is a GET request handler /api/nextdate?now=<now>&date=<date>&repeat=<repeat>, here:
//   - <now> is today's date
//   - <date> - the date of the last task run
//   - <repeat> - task repetition rules
//
// It receives and validates GET parameters and returns the date of the next task execution if successful, otherwise it returns an error.
func NextDayHandler(w http.ResponseWriter, r *http.Request) {
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
