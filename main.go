package main

import (
	"context"
	"git-learn/postgres/db_connection"
	"git-learn/postgres/sql"
)

func main() {

	ctx := context.Background()

	conn, err := db_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	sql.CreateTable(ctx, conn)

}
