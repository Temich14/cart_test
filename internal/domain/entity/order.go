package entity

import (
	"gorm.io/gorm"
	"time"
)

type OrderStatus string

const (
	CREATED     OrderStatus = "created"
	IN_PROGRESS OrderStatus = "in_progress"
	CANCELED    OrderStatus = "cancelled"
	COMPLETED   OrderStatus = "completed"
)

type Order struct {
	gorm.Model
	UserID        uint
	Cost          float32
	RawCost       float32
	ItemsQuantity int
	Status        string
	CreatedAt     time.Time
	CompletedAt   time.Time
	Items         []OrderItem
}
type OrderItem struct {
	gorm.Model
	OrderID   uint
	ProductID uint
	Quantity  int
	Cost      float32
	RawCost   float32
}
