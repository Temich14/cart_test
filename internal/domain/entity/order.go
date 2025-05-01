package entity

import "time"

type Order struct {
	ID        string `gorm:"primaryKey"`
	UserID    string
	Status    string
	CreatedAt time.Time
	Items     []OrderItem
}
type OrderItem struct {
	ID        string `gorm:"primaryKey"`
	OrderID   string
	ProductID string
	Quantity  int
}
