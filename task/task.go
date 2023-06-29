package task

import (
	"database/sql"
	"time"
)

type Priority string

const (
	Low    Priority = "low"
	Medium Priority = "medium"
	High   Priority = "high"
)

type Task struct {
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	DeletedAt   sql.NullTime `json:"deletedAt" gorm:"index"`
	ID          uint         `json:"id" gorm:"primaryKey"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	DueDate     string       `json:"dueDate"`
	Priority    Priority     `json:"priority"`
	Completed   bool         `json:"completed"`
	UserID      uint         `json:"userID"`
}
