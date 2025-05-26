package api

import (
	"net/http"
)

func Init() {
	http.HandleFunc("/api/nextdate", nextDayHandler)
	http.HandleFunc("/api/task", AuthMiddleware(taskHandler))
	http.HandleFunc("/api/tasks", AuthMiddleware(searchHandler))
	http.HandleFunc("/api/task/done", AuthMiddleware(doneHandler))
	http.HandleFunc("/api/signin", loginHandler)
}
