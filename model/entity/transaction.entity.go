package entity

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Amount    float64        `json:"amount"`
	UserID    uint           `json:"user_id"`
	User      User           `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
