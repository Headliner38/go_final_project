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
		/*case http.MethodGet:
			getTaskHandler(res, req)
		case http.MethodDelete:
			deleteTaskHandler(res, req)
		case http.MethodPut:
			putTaskHandler(res, req)*/
	}
}

func writeJson(res http.ResponseWriter, statusCode int, data any) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(res).Encode(data)
	/*jsonBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	_, err = res.Write(jsonBytes)
	if err != nil {
		panic(err)
	}*/
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
