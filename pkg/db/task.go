package db

import (
	"database/sql"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func AddTask(task *Task) (int64, error) {
	Init()
	defer database.Close()

	res, err := database.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)",
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat))
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func GetTasks(limit int) ([]*Task, error) {
	Init()
	defer database.Close()

	rows, err := database.Query("SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date LIMIT :limit",
		sql.Named("limit", limit))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*Task
	for rows.Next() {
		t := &Task{}

		err := rows.Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat)
		if err != nil {
			return nil, err
		}

		res = append(res, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	if res == nil {
		return make([]*Task, 0), nil
	}

	return res, nil
}
