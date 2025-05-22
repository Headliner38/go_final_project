package db

import (
	"database/sql"
	"fmt"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

var DB *sql.DB

func AddTask(task *Task) (int64, error) {
	var id int64

	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?) returning id;`
	rows, err := DB.Query(query, task.Date, task.Title, task.Comment, task.Repeat)
	fmt.Print(DB)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return 0, fmt.Errorf("не удалось получить id: %w", err)
		}
	}
	return id, nil
}

func Tasks(limit int) ([]*Task, error) {
	// мы делаем селект
	// Задачи должны быть отсортированы по дате в сторону увеличения. Каждая задача должна содержать все поля таблицы scheduler в виде строк
	query := `SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date ASC LIMIT ?`

	rows, err := DB.Query(query, limit)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании запроса: %v", err)
	}
	defer rows.Close()
	var tasks []*Task
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat)
		if err != nil {
			return nil, fmt.Errorf("ошибка чтения строк: %v", err)
		}
		tasks = append(tasks, &t)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка: %v", err)
	}

	return tasks, nil
}

func GetTask(id string) (*Task, error) {
	query := `SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?`
	row := DB.QueryRow(query, id)
	var task Task
	err := row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении: %v", err)
	}
	return &task, nil
}

func UpdateTask(task *Task) error {
	query := `UPDATE scheduler SET 
		date = ?,
		title = ?,
		comment = ?,
		repeat = ?
		WHERE id = ?	
	`
	result, err := DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat, task.ID)
	if err != nil {
		return fmt.Errorf("ошибка выполнения запроса: %v", err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка: %v", err)
	}
	if count == 0 {
		return fmt.Errorf(`неверный id для обновления задачи`)
	}
	return nil
}

func DeleteTask(id string) error {
	query := `DELETE FROM scheduler WHERE id = ?`

	result, err := DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("невозможно выполнить запрос: %v", err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка: %v", err)
	}
	if count == 0 {
		return fmt.Errorf(`неверный id для удаления задачи`)
	}
	return nil
}

func UpdateDate(next string, id string) error {
	query := `UPDATE scheduler SET
		date = ?
		WHERE id = ?
	`
	result, err := DB.Exec(query, next, id)
	if err != nil {
		return fmt.Errorf("ошибка выполения запроса: %v", err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка: %v", err)
	}
	if count == 0 {
		return fmt.Errorf(`неверный id`)
	}
	return nil
}
