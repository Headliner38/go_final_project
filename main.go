package main

import (
	"fmt"
	"log"

	"go_final_project/pkg/db"
	"go_final_project/pkg/server"
)

func main() {
	conn, err := db.Init("scheduler.db")
	if err != nil {
		log.Fatalf("Ошибка открытия базы данных: %v", err)
	}
	db.DB = conn

	err = server.Run()
	if err != nil {
		fmt.Print("Сервер не запускается")
	}
}
