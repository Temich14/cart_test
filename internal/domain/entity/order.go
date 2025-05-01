package entity

import "time"

type Order struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Status    string
	CreatedAt time.Time
	Items     []OrderItem
}
type OrderItem struct {
	ID        string `gorm:"primaryKey"`
	OrderID   uint
	ProductID uint
	Quantity  int
}
