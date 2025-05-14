package order

import "github.com/Temich14/cart_test/internal/domain/entity"

type OrderService interface {
	CreateNewOrder(userID uint) (*entity.Order, error)
	ChangeStatus(orderID uint, status entity.OrderStatus) (*entity.Order, error)
	GetOrders(userID uint, status string, page, limit int) (*entity.OrderPaginationResponse, error)
	GetOrder(orderID uint) (*entity.Order, error)
}
