package order

import "github.com/Temich14/cart_test/internal/domain/entity"

type Service struct {
	repo Repository
}

func NewOrderService(repo Repository) *Service {
	return &Service{}
}

func (s *Service) CreateNewOrder(cart *entity.Cart) (*entity.Order, error) {
	order, err := s.repo.CreateOrder(cart)
	if err != nil {
		return nil, err
	}
	return order, nil
}
func (s *Service) ChangeStatus(orderID uint, status entity.OrderStatus) error {
	err := s.repo.ChangeOrderStatus(orderID, status)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) GetOrders(userID uint) ([]*entity.Order, error) {
	orders, err := s.repo.GetUserOrders(userID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
func (s *Service) GetOrder(orderID uint) (*entity.Order, error) {
	order, err := s.repo.GetUserOrder(orderID)
	if err != nil {
		return nil, err
	}
	return order, nil
}
