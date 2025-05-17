package api

import (
	"fmt"
	"go_final_project/pkg/db"
	"time"
)

func checkDate(task *db.Task) error {
	now := time.Now()

	if task.Date == "" {
		task.Date = now.Format("20060102")
	}

	t, err := time.Parse("20060102", task.Date)
	if err != nil {
		return fmt.Errorf("неверный формат даты: %v", err)
	}

	if task.Repeat != "" {
		next, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return fmt.Errorf("неверное правило повторения")
		}
		if afterNow(now, t) {
			if len(task.Repeat) == 0 {
				// если правила повторения нет, то берём сегодняшнее число
				task.Date = now.Format("20060102")
			} else {
				// в противном случае, берём вычисленную ранее следующую дату
				task.Date = next
			}
		}
	}

	return nil
}
