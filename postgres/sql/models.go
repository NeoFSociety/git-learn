package sql

import "time"

type Task struct {
	Id           int
	Title        string
	Description  string
	Completed    bool
	Created_at   time.Time
	Completed_at *time.Time
}
