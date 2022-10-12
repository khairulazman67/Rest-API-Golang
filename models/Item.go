package models

import (
	"time"
)

type Item struct {
	ID          uint   `json:"ID" form:"ID" gorm:"primaryKey"`
	Item_code   uint   `json:"Item_code" form:"Item_code" gorm:"not null"`
	Description string `json:"Description" gorm:"not null;type:varchar(191)"`
	Quantity    uint   `json:"Quantity"`
	Order_id    uint   `json:"Order_id" gorm:"not null"`
	CretaedAt   time.Time
	UpdatedAt   time.Time
}
