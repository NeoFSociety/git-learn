package main

import (
	"context"
	"fmt"
	"git-learn/postgres/db_connection"
	"git-learn/postgres/sql"
)

func main() {

	ctx := context.Background()

	conn, err := db_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	tasks, err := sql.GetData(ctx, conn)

	if err != nil {
		panic(err)
	}

	for _, task := range tasks {
		fmt.Println("-----------------------------")
		fmt.Println("id:", task.Id)
		fmt.Println("title", task.Title)
		fmt.Println("description:", task.Description)
		fmt.Println("completed", task.Completed)
		fmt.Println("created_at:", task.Created_at)
		fmt.Println("completed_at", task.Completed_at)
	}

	fmt.Println("SUCCESS")
}
