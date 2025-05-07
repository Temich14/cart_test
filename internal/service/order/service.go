package order

import (
	"log/slog"
	"runtime/debug"

	"github.com/Temich14/cart_test/internal/domain/entity"
)

type Service struct {
	repo Repository
	log  *slog.Logger
}

func NewOrderService(repo Repository, log *slog.Logger) *Service {
	return &Service{repo: repo, log: log}
}

func (s *Service) CreateNewOrder(userID uint) (*entity.Order, error) {
	s.log.Debug("creating new order", slog.Uint64("user_id", uint64(userID)))

	order, err := s.repo.CreateOrder(userID)
	if err != nil {
		s.log.Error(
			"failed to create order",
			slog.Uint64("user_id", uint64(userID)),
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return nil, err
	}

	s.log.Debug("order created", slog.Uint64("user_id", uint64(userID)), slog.Uint64("order_id", uint64(order.ID)))
	return order, nil
}

func (s *Service) ChangeStatus(orderID uint, status entity.OrderStatus) (*entity.Order, error) {
	s.log.Debug("changing order status", slog.Uint64("order_id", uint64(orderID)), slog.String("new_status", string(status)))

	order, err := s.repo.ChangeOrderStatus(orderID, status)
	if err != nil {
		s.log.Error(
			"failed to change order status",
			slog.Uint64("order_id", uint64(orderID)),
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return nil, err
	}

	s.log.Debug("order status changed", slog.Uint64("order_id", uint64(order.ID)), slog.String("new_status", string(order.Status)))
	return order, nil
}

func (s *Service) GetOrders(userID uint, status string, page, limit int) (*entity.OrderPaginationResponse, error) {
	s.log.Debug("retrieving user orders",
		slog.Uint64("user_id", uint64(userID)),
		slog.String("status", status),
		slog.Int("page", page),
		slog.Int("limit", limit),
	)

	orders, err := s.repo.GetUserOrders(userID, status, page, limit)
	if err != nil {
		s.log.Error(
			"failed to retrieve user orders",
			slog.Uint64("user_id", uint64(userID)),
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return nil, err
	}

	s.log.Debug("user orders retrieved", slog.Uint64("user_id", uint64(userID)))
	return orders, nil
}

func (s *Service) GetOrder(orderID uint) (*entity.Order, error) {
	s.log.Debug("retrieving order", slog.Uint64("order_id", uint64(orderID)))

	order, err := s.repo.GetUserOrder(orderID)
	if err != nil {
		s.log.Error(
			"failed to retrieve order",
			slog.Uint64("order_id", uint64(orderID)),
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return nil, err
	}

	s.log.Debug("order retrieved", slog.Uint64("order_id", uint64(orderID)))
	return order, nil
}
