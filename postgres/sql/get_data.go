package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func GetData(ctx context.Context, conn *pgx.Conn) ([]Task, error) {

	data := []Task{}

	sqlQuery := `
	SELECT id, title, description, completed, created_at, completed_at
	FROM tasks
	`

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {

		var task Task

		err := rows.Scan(
			&task.Id,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.Created_at,
			&task.Completed_at,
		)

		if err != nil {
			return nil, err
		}

		data = append(data, task)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return data, err
}
