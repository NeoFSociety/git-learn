package db_connection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	fmt.Println("Подключение к БД успешно установлено")

	return pgx.Connect(ctx, "postgres://postgres:777@localhost:5432/postgres")
}
