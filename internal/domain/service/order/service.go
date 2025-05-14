// Package order содержит реализацию бизнес-логики для работы с заказами пользователей.
package order

import (
	"github.com/Temich14/cart_test/internal/domain/service"
	"log/slog"
	"runtime/debug"

	"github.com/Temich14/cart_test/internal/domain/entity"
)

type Service struct {
	repo            Repository
	log             *slog.Logger
	productProvider service.ProductProvider
}

// NewOrderService создаёт новый экземпляр сервиса заказов.
//   - repo: интерфейс репозитория для доступа к данным заказов.
//   - log: логгер для записи событий.
//   - prov: провайдер продуктов для подгрузки информации о товарах.
func NewOrderService(repo Repository, log *slog.Logger, prov service.ProductProvider) *Service {
	return &Service{repo: repo, log: log, productProvider: prov}
}

// CreateNewOrder создаёт новый заказ для указанного пользователя.
// Принимает:
//   - userID: идентификатор пользователя.
//
// Возвращает:
//   - *entity.Order: созданный заказ.
//   - error: ошибка, если заказ не удалось создать.
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

// ChangeStatus изменяет статус существующего заказа.
//
// Принимает:
//   - orderID: идентификатор заказа.
//   - status: новый статус заказа.
//
// Возвращает:
//   - *entity.Order: заказ с обновлённым статусом.
//   - error: ошибка, если изменение не удалось.
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

// GetOrders возвращает список заказов пользователя с пагинацией и фильтрацией по статусу.
//
// Принимает:
//   - userID: идентификатор пользователя.
//   - status: строковое представление статуса заказа (опционально).
//   - page: номер страницы.
//   - limit: количество заказов на странице.
//
// Возвращает:
//   - *entity.OrderPaginationResponse: структура с данными о заказах и пагинацией.
//   - error: ошибка, если заказы не удалось получить.
func (s *Service) GetOrders(userID uint, status string, page, limit int) (*entity.OrderPaginationResponse, error) {
	s.log.Debug("retrieving user orders",
		slog.Uint64("user_id", uint64(userID)),
		slog.String("status", status),
		slog.Int("page", page),
		slog.Int("limit", limit),
	)

	orders, err := s.repo.GetUserOrders(userID, status, page, limit)
	if err != nil {
		s.log.Error("failed to retrieve user orders", slog.String("error", err.Error()))
		return nil, err
	}

	var productIDs []uint
	for _, order := range orders.Data {
		for _, item := range order.Items {
			productIDs = append(productIDs, item.ProductID)
		}
	}

	productsMap, err := s.productProvider.GetProductsByIDs(productIDs)
	if err != nil {
		s.log.Error("failed to fetch products", slog.String("error", err.Error()))
		return nil, err
	}

	for _, order := range orders.Data {
		for i := range order.Items {
			item := &order.Items[i]
			if p, ok := productsMap[item.ProductID]; ok {
				item.Product = *p
			} else {
				s.log.Warn("missing product", slog.Uint64("product_id", uint64(item.ProductID)))
			}
		}
	}

	s.log.Debug("user orders retrieved", slog.Uint64("user_id", uint64(userID)))
	return orders, nil
}

// GetOrder возвращает заказ по его идентификатору.
//
// Принимает:
//   - orderID: идентификатор заказа.
//
// Возвращает:
//   - *entity.Order: найденный заказ.
//   - error: ошибка, если заказ не найден или произошла ошибка при получении.
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

	var productIDs []uint
	for _, item := range order.Items {
		productIDs = append(productIDs, item.ProductID)
	}

	productsMap, err := s.productProvider.GetProductsByIDs(productIDs)
	if err != nil {
		s.log.Error("failed to fetch products for order", slog.String("error", err.Error()))
		return nil, err
	}

	for i := range order.Items {
		item := &order.Items[i]
		if p, ok := productsMap[item.ProductID]; ok {
			item.Product = *p
		} else {
			s.log.Warn("missing product for order item", slog.Uint64("product_id", uint64(item.ProductID)))
		}
	}

	s.log.Debug("order retrieved", slog.Uint64("order_id", uint64(orderID)))
	return order, nil
}
