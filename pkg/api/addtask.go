package api

import (
	"encoding/json"
	"fmt"
	"go_final_project/pkg/db"
	"net/http"
)

func taskHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		addTaskHandler(res, req)
	case http.MethodGet:
		getTaskHandler(res, req)
	case http.MethodPut:
		putTaskHandler(res, req)
	case http.MethodDelete:
		deleteTaskHandler(res, req)
	default:
		http.Error(res, "данный метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

func getTaskHandler(res http.ResponseWriter, req *http.Request) {
	idStr := req.URL.Query().Get("id")
	if len(idStr) == 0 {
		writeJson(res, http.StatusBadRequest, map[string]string{"error": "не указан id"})
		return
	}
	needTask, err := db.GetTask(idStr)
	if err != nil {
		writeJson(res, http.StatusBadRequest, map[string]string{"error": "ошибка получения задания по id" + err.Error()})
		return
	}
	writeJson(res, http.StatusOK, needTask)
}

func putTaskHandler(res http.ResponseWriter, req *http.Request) {
	var task db.Task
	err := json.NewDecoder(req.Body).Decode(&task)
	if err != nil {
		writeJson(res, http.StatusBadRequest, map[string]string{"error": "ошибка json: " + err.Error()})
		return
	}

	if task.Title == "" {
		writeJson(res, http.StatusBadRequest, map[string]string{"error": "пустой заголовок задачи"})
		return
	}

	err = checkDate(&task)
	if err != nil {
		writeJson(res, http.StatusBadRequest, map[string]string{"error": "ошибка checkDate: " + err.Error()})
		return
	}

	err = db.UpdateTask(&task)
	if err != nil {
		writeJson(res, http.StatusInternalServerError, map[string]string{"error": "ошибка AddTask: " + err.Error()})
		return
	}
	writeJson(res, http.StatusOK, map[string]string{})
}

func deleteTaskHandler(res http.ResponseWriter, req *http.Request) {
	idStr := req.URL.Query().Get("id")
	if len(idStr) == 0 {
		writeJson(res, http.StatusBadRequest, map[string]string{"error": "ошибка получения id"})
		return
	}
	err := db.DeleteTask(idStr)
	if err != nil {
		writeJson(res, http.StatusBadRequest, map[string]string{"error": "ошибка удаления задачи"})
		return
	}
	writeJson(res, http.StatusOK, map[string]string{})
}

func writeJson(res http.ResponseWriter, statusCode int, data any) {
	_ = statusCode
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := json.NewEncoder(res).Encode(data)
	if err != nil {
		fmt.Printf("ошибка encode: %v\n", err)
	}

}

func addTaskHandler(res http.ResponseWriter, req *http.Request) {
	var task db.Task
	err := json.NewDecoder(req.Body).Decode(&task)
	if err != nil {
		writeJson(res, http.StatusBadRequest, map[string]string{"error": "ошибка json: " + err.Error()})
		return
	}

	if task.Title == "" {
		writeJson(res, http.StatusBadRequest, map[string]string{"error": "пустой заголовок задачи"})
		return
	}

	err = checkDate(&task)
	if err != nil {
		writeJson(res, http.StatusBadRequest, map[string]string{"error": "ошибка checkDate: " + err.Error()})
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		writeJson(res, http.StatusInternalServerError, map[string]string{"error": "ошибка AddTask: " + err.Error()})
		return
	}
	writeJson(res, http.StatusOK, map[string]string{"id": fmt.Sprintf("%d", id)})
}
