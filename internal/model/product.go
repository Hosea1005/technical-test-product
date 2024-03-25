package model

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID        uuid.UUID  `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"type:varchar(100)" json:"name"`
	Type      string     `gorm:"type:varchar(20)" json:"type"`
	Price     float64    `gorm:"type:float" json:"price"`
	CreatedAt *time.Time ` json:"created_at"`
	UpdatedAt *time.Time ` json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
