package model

import (
	"time"
)

//产品
type Products struct {
    //
    Id int  `gorm:"column:id"`
    //
    CreatedAt time.Time  `gorm:"column:created_at"`
    //
    UpdatedAt time.Time  `gorm:"column:updated_at"`
    //
    DeletedAt string  `gorm:"column:deleted_at"`
    //
    Code string  `gorm:"column:code"`
    //
    Price int  `gorm:"column:price"`
    
}

//产品表名
func (Products) TableName() string {
	return "products"
}
