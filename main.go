package main

import (
	"context"
	"fmt"
	"git-learn/postgres/db_connection"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Файл .env не найден, используем системные переменные")
	}

	ctx := context.Background()

	_, err := db_connection.CreateConnection(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println("SUCCESS")
}
