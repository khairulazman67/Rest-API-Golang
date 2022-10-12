package models

import (
	"time"
)

type Order struct {
	ID            uint   `gorm:"primaryKey"`
	Customer_name string `gorm:"not null; type:varchar(191)"`
	Ordered_at    time.Time
	UpdatedAt     time.Time
}
