package server

import (
	"fmt"
	"go_final_project/pkg/api"
	"net/http"
	"os"
	"strconv"
)

const defaultPort = 7540

func checkPort() int { // задание со звездочкой для TODO_PORT
	portStr := os.Getenv("TODO_PORT")
	if len(portStr) == 0 {
		fmt.Printf("порт из переменной окружения не задан, используется порт по умолчанию\n")
		return defaultPort
	}

	port, err := strconv.Atoi(portStr)
	if err != nil || port < 1 || port > 65535 {
		return defaultPort
	}
	return port

}

func Run() error {
	port := checkPort()
	api.Init()
	http.Handle("/", http.FileServer(http.Dir("web")))
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
