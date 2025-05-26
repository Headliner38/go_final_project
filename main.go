package main

import (
	"fmt"
	"log"

	"go_final_project/pkg/api/auth"
	"go_final_project/pkg/db"
	"go_final_project/pkg/server"
)

func main() {
	err := auth.InitKeys()
	if err != nil {
		log.Fatalf("Ошибка инициализации RSA ключей: %v", err)
	}

	conn, err := db.Init()
	if err != nil {
		log.Fatalf("Ошибка открытия базы данных: %v", err)
	}
	db.DB = conn

	err = server.Run()
	if err != nil {
		fmt.Print("Сервер не запускается")
	}
}
