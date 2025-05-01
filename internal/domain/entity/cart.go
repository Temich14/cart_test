package entity

import "time"

type Cart struct {
	ID        string `gorm:"primary_key"`
	UserID    string
	items     []CartItem
	createdAt time.Time
	updatedAt time.Time
}
type CartItem struct {
	ID        string `gorm:"primary_key"`
	ProductID string
	Quantity  int
	createdAt time.Time
}
