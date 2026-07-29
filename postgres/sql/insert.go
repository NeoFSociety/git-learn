package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func AddTask(ctx context.Context, conn *pgx.Conn, task Task) error {

	sqlQuery := `
	INSERT INTO tasks (title, description, completed, created_at, completed_at)
	VALUES ($1, $2, $3, $4, $5)
	`

	_, err := conn.Exec(
		ctx,
		sqlQuery,
		task.Title,
		task.Description,
		task.Completed,
		task.Created_at,
		task.Completed_at,
	)

	return err
}
