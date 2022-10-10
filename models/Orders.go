package models

import (
	"time"
)

type Orders struct {
	ID            uint   `gorm:"primaryKey"`
	Customer_name string `gorm:"not null; type:varchar(191)"`
	Ordered_at    time.Time
	UpdatedAt     time.Time
}

// func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
// 	fmt.Println("Product Before Create()")
// 	if len(p.Name) < 4 {
// 		err = errors.New("Product name is too short")
// 	}
// 	return
// }
