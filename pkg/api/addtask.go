package api

import (
	"encoding/json"
	"fmt"
	"go_final_project/pkg/api/auth"
	"go_final_project/pkg/db"
	"net/http"
	"os"
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

func WriteJSON(res http.ResponseWriter, statusCode int, data any) {
	writeJson(res, statusCode, data)
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

type AuthRequest struct {
	Password string `json:"password"`
}

func loginHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		writeJson(res, http.StatusMethodNotAllowed, map[string]string{"error": "неподдерживаемый метод"})
		return
	}

	correctPass := os.Getenv("TODO_PASSWORD")
	if len(correctPass) == 0 {
		writeJson(res, http.StatusBadRequest, map[string]string{"error": "пароль в переменной окружения не задан"})
		return
	}

	var au AuthRequest
	err := json.NewDecoder(req.Body).Decode(&au)
	if err != nil {
		writeJson(res, http.StatusBadRequest, map[string]string{"error": "неверный запрос"})
		return
	}

	//fmt.Println("my pass: ", passStr, "cor pas: ", correctPass) //смотрим пароли для себя
	if au.Password != correctPass {
		writeJson(res, http.StatusUnauthorized, map[string]string{"error": "Неверный пароль"})
		return
	}

	token, err := auth.GenerateToken(correctPass) //создаем токен
	if err != nil {
		writeJson(res, http.StatusUnauthorized, map[string]string{"error": "Не удалось сгенерировать jwt token: " + err.Error()})
		return
	}
	writeJson(res, http.StatusOK, map[string]string{"token": token})
}
