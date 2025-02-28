package entities

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Name      string         `json:"name" gorm:"not null"`
	Price     int            `json:"price" gorm:"unique"`
}













