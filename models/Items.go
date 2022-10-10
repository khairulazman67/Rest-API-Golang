package models

import (
	"time"
)

type Items struct {
	ID          uint   `gorm:"primaryKey"`
	Item_code   uint   `gorm:"not null"`
	Description string `gorm:"not null;type:varchar(191)"`
	Quantity    uint
	CretaedAt   time.Time
	UpdatedAt   time.Time
}
