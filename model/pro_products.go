package model

import (
	"time"
)

//商品
type ProProducts struct {
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

//商品表名
func (ProProducts) TableName() string {
	return "pro_products"
}
