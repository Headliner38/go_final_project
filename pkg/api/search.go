package api

import (
	"go_final_project/pkg/db"
	"net/http"
	"time"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func searchHandler(res http.ResponseWriter, req *http.Request) {

	searchStr := req.URL.Query().Get("search")
	if searchStr == "" {
		tasks, err := db.Tasks(50) // в параметре максимальное количество записей
		if err != nil {
			writeJson(res, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		writeJson(res, http.StatusOK, TasksResp{Tasks: tasks})
		return
	}
	if date, err := time.Parse("02.01.2006", searchStr); err == nil {
		tasks, err := db.SearchTaskByDate(date.Format("20060102"), 50)
		if err != nil {
			writeJson(res, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		writeJson(res, http.StatusOK, TasksResp{Tasks: tasks})
		return
	} else {
		tasks, err := db.SearchTaskByWord(searchStr, 50)
		if err != nil {
			writeJson(res, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		writeJson(res, http.StatusOK, TasksResp{Tasks: tasks})
		return
	}
}
