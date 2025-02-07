package models

import (
	"time"
)

type Task struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"` // 🔹 Hubungkan dengan user
	Text      string    `json:"text"`
	Date      string    `json:"date"` // Format "YYYY-MM-DD"
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
