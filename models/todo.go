package models

import "gorm.io/gorm"

// Todo struct untuk tabel todo dalam database
type Todo struct {
	gorm.Model
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}
