package main

import (
	"fmt"

	"go_final_project/pkg/db"
	"go_final_project/pkg/server"
)

func main() {
	err := db.Init("scheduler.db")
	if err != nil {
		fmt.Printf("Ошибка создания БД: %s", err)
	}

	err = server.Run()
	if err != nil {
		fmt.Print("Сервер не запускается")
	}
}
