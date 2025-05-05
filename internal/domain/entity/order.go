package entity

import (
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
	ID            uint `gorm:"primarykey" json:"id" example:"1"`
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
	ID        uint `gorm:"primarykey" json:"id" example:"1"`
	OrderID   uint
	ProductID uint
	Quantity  int
	Cost      float32
	RawCost   float32
}
type OrderPaginationResponse struct {
	Data []*Order       `json:"data"`
	Meta PaginationMeta `json:"meta"`
}
