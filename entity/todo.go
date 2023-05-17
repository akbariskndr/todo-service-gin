package entity

import "time"

type Todo struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Title     string     `json:"title"`
	Completed bool       `json:"completed"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}