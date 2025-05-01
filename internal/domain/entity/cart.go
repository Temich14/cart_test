package entity

import (
	"gorm.io/gorm"
	"time"
)

type Cart struct {
	gorm.Model
	UserID        uint
	TotalQuantity int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Items         []CartItem
}
type CartItem struct {
	gorm.Model
	CartID    uint
	ProductID uint
	Quantity  int
	CreatedAt time.Time
}
