package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Todos    []Todo `gorm:"foreignKey:UserID" json:"todos"` // Relasi ke Todo
}
