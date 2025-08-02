package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID   uint `json:"user_id" gorm:"not null"`
	ItemID   uint `json:"item_id" gorm:"not null"`
	Quantity int  `json:"quantity" gorm:"default:1"`
	User     User `json:"user" gorm:"foreignKey:UserID"`
	Item     Item `json:"item" gorm:"foreignKey:ItemID"`
}