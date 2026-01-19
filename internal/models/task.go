package models

import "time"

type Task struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"not null"`
	Status    string    `gorm:"default:'todo'"`
	UserID    uint      `gorm:"index;not null"`
	CreatedAt time.Time 
	UpdatedAt time.Time 
}
