package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"size:255"`
	Email     string `gorm:"type:varchar(100);unique_index"`
	Password  string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Product struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"size:255"`
	Description string `gorm:"size:1000"`
	Price       float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
