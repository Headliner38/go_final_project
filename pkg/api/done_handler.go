package api

import (
	"go_final_project/pkg/db"
	"net/http"
	"time"
)

func doneHandler(res http.ResponseWriter, req *http.Request) {
	idStr := req.URL.Query().Get("id")
	needTask, err := db.GetTask(idStr)
	if err != nil {
		writeJson(res, http.StatusBadRequest, map[string]string{"error": "ошибка получения задачи"})
		return
	}

	if len(needTask.Repeat) == 0 {
		err := db.DeleteTask(idStr)
		if err != nil {
			writeJson(res, http.StatusBadRequest, map[string]string{"error": "ошибка удаления задачи"})
			return
		}
	} else {
		now := time.Now()
		str, err := NextDate(now, needTask.Date, needTask.Repeat)
		if err != nil {
			writeJson(res, http.StatusBadRequest, map[string]string{"error": "ошибка получения следующей даты"})
			return
		}
		db.UpdateDate(str, idStr)
	}
	writeJson(res, http.StatusOK, map[string]string{})
}
