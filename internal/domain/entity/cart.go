package entity

import (
	"gorm.io/gorm"
	"time"
)

type Cart struct {
	gorm.Model
	UserID        uint
	TotalQuantity int
	createdAt     time.Time
	updatedAt     time.Time
	items         []CartItem
}
type CartItem struct {
	gorm.Model
	CartID    uint
	ProductID uint
	Quantity  int
	createdAt time.Time
}
