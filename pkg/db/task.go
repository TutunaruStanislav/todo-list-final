package db

import (
	"database/sql"
	"errors"
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

func GetTask(id int64) (*Task, error) {
	Init()
	defer database.Close()

	task := &Task{}

	row := database.QueryRow("SELECT id, date, title, comment, repeat FROM scheduler WHERE id = :id", sql.Named("id", id))
	err := row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return nil, errors.New("task not found")
	}

	return task, nil
}

func UpdateTask(task *Task) error {
	Init()
	defer database.Close()

	res, err := database.Exec("UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id",
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
		sql.Named("id", task.ID))

	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("incorrect id for updating task")
	}
	return nil
}

func DeleteTask(id int64) error {
	Init()
	defer database.Close()

	res, err := database.Exec("DELETE FROM scheduler WHERE id = :id", sql.Named("id", id))
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("there were no rows deleted")
	}

	return nil
}
