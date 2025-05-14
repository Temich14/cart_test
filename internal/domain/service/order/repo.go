package order

import "github.com/Temich14/cart_test/internal/domain/entity"

// Repository интерфейс для взаимодействия с бд.
type Repository interface {
	CreateOrder(userID uint) (*entity.Order, error)
	ChangeOrderStatus(orderID uint, status entity.OrderStatus) (*entity.Order, error)
	GetUserOrders(userID uint, status string, page, limit int) (*entity.OrderPaginationResponse, error)
	GetUserOrder(orderID uint) (*entity.Order, error)
}
