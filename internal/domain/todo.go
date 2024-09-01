package domain

import (
	"database/sql"
	"time"
)

type TodoItem struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	DueDate     time.Time      `json:"due_date"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
}
