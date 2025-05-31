package db

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// a
func SearchTaskByWord(search string, limit int) ([]*Task, error) {
	tasks := make([]*Task, 0)
	searchPtrn := "%" + search + "%"
	query := ` SELECT id, date, title, comment, repeat FROM scheduler WHERE title LIKE ? OR comment LIKE ? ORDER BY date LIMIT ?`
	rows, err := DB.Query(query, searchPtrn, searchPtrn, limit)
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		task := new(Task)
		err = rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, fmt.Errorf("ошибка rows.scan %v", err)
		}
		tasks = append(tasks, task)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("ошибка rows.err %v", err)
	}
	return tasks, nil
}

func SearchTaskByDate(search string, limit int) ([]*Task, error) {
	tasks := make([]*Task, 0)
	query := `SELECT id, date, title, comment, repeat FROM scheduler WHERE date = ? LIMIT ?`
	rows, err := DB.Query(query, search, limit)
	if err != nil {
		return nil, fmt.Errorf("ошибка при запросе %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		task := new(Task)
		err = rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, fmt.Errorf("ошибка rows.scan %v", err)
		}
		tasks = append(tasks, task)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("о rows.err %v", err)
	}
	return tasks, nil
}
