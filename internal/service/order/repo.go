package order

import "github.com/Temich14/cart_test/internal/domain/entity"

type Repository interface {
	CreateOrder(cart *entity.Cart) (*entity.Order, error)
	ChangeOrderStatus(orderID uint, status entity.OrderStatus) error
	GetUserOrders(userID uint) ([]*entity.Order, error)
	GetUserOrder(orderID uint) (*entity.Order, error)
}
