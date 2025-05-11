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
	ID            uint    `gorm:"primarykey" json:"id" example:"1"`
	UserID        uint    `json:"user_id" example:"1"`
	Cost          float32 `json:"total_cost"`
	ItemsQuantity int     `json:"total_quantity"`
	Status        string
	CreatedAt     time.Time
	CompletedAt   time.Time
	Items         []OrderItem
}
type OrderItem struct {
	ID        uint `gorm:"primarykey" json:"id" example:"1"`
	OrderID   uint
	ProductID uint
	Product   Product
	Quantity  int
	Cost      float32
}
type OrderPaginationResponse struct {
	Data []*Order       `json:"data"`
	Meta PaginationMeta `json:"meta"`
}
