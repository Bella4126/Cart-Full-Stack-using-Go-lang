package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID uint  `json:"user_id" gorm:"not null"`
	Total  float64 `json:"total" gorm:"not null"`
	User   User  `json:"user" gorm:"foreignKey:UserID"`
}

type OrderItem struct {
	gorm.Model
	OrderID  uint    `json:"order_id" gorm:"not null"`
	ItemID   uint    `json:"item_id" gorm:"not null"`
	Quantity int     `json:"quantity" gorm:"not null"`
	Price    float64 `json:"price" gorm:"not null"`
	Order    Order   `json:"order" gorm:"foreignKey:OrderID"`
	Item     Item    `json:"item" gorm:"foreignKey:ItemID"`
}