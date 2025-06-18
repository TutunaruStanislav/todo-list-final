package db

import (
	"database/sql"
	"errors"
)

// Model for task
type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

// AddTask - adds a record with a new task to DB.
func AddTask(db *sql.DB, task *Task) (int64, error) {
	res, err := db.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)",
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat))
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// GetTasks - finds a list of tasks in DB by specified parameters.
func GetTasks(db *sql.DB, limit int, search string, date string) ([]*Task, error) {
	var rows *sql.Rows
	var err error
	if len(search) > 0 {
		search := "%" + search + "%"
		rows, err = db.Query("SELECT * FROM scheduler WHERE LOWER(title) LIKE LOWER(:search) OR LOWER(comment) LIKE LOWER(:search) ORDER BY date LIMIT :limit",
			sql.Named("limit", limit),
			sql.Named("search", search))
	} else if len(date) > 0 {
		rows, err = db.Query("SELECT * FROM scheduler WHERE date = :date LIMIT :limit",
			sql.Named("limit", limit),
			sql.Named("date", date))
	} else {
		rows, err = db.Query("SELECT * FROM scheduler ORDER BY date LIMIT :limit",
			sql.Named("limit", limit))
	}

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

// GetTasks - finds task by id in DB.
func GetTask(db *sql.DB, id int64) (*Task, error) {
	task := &Task{}

	row := db.QueryRow("SELECT id, date, title, comment, repeat FROM scheduler WHERE id = :id", sql.Named("id", id))
	err := row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return nil, errors.New("task not found")
	}

	return task, nil
}

// UpdateTask- updates the DB task by id.
func UpdateTask(db *sql.DB, task *Task) error {
	res, err := db.Exec("UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id",
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

// DeleteTask - deletes task by id in DB.
func DeleteTask(db *sql.DB, id int64) error {
	res, err := db.Exec("DELETE FROM scheduler WHERE id = :id", sql.Named("id", id))
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return sql.ErrNoRows
	}

	return nil
}
