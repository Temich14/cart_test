package order

import "github.com/Temich14/cart_test/internal/domain/entity"

type Service struct {
	repo Repository
}

func NewOrderService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateNewOrder(userID uint) (*entity.Order, error) {
	order, err := s.repo.CreateOrder(userID)
	if err != nil {
		return nil, err
	}
	return order, nil
}
func (s *Service) ChangeStatus(orderID uint, status entity.OrderStatus) (*entity.Order, error) {
	order, err := s.repo.ChangeOrderStatus(orderID, status)
	if err != nil {
		return nil, err
	}
	return order, nil
}
func (s *Service) GetOrders(userID uint, status string, page, limit int) (*entity.OrderPaginationResponse, error) {
	orders, err := s.repo.GetUserOrders(userID, status, page, limit)
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
