package entity

import (
	"time"
)

type Cart struct {
	ID            uint `gorm:"primarykey" json:"id" example:"1"`
	UserID        uint
	TotalQuantity int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Items         []CartItem
}
type CartItem struct {
	ID        uint      `gorm:"primarykey" json:"id" example:"1"`
	CartID    uint      `json:"cart_id" example:"1"`
	ProductID uint      `json:"product_id" example:"1"`
	Quantity  int       `json:"quantity" example:"1"`
	CreatedAt time.Time `json:"created_at" example:"2025-05-05 10:30:00"`
}
type PaginationMeta struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}
type CartWithItemsPagination struct {
	ID            uint           `json:"id"`
	UserID        uint           `json:"user_id"`
	TotalQuantity int            `json:"total_quantity"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	Items         []CartItem     `json:"items"`
	ItemsMeta     PaginationMeta `json:"items_meta"`
}
