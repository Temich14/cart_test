package order

import (
	"github.com/Temich14/cart_test/internal/domain/entity"
	"github.com/stretchr/testify/mock"
	"io"
	"log/slog"
)

type MockRepo struct {
	mock.Mock
}

func NewServiceWithMock() (*Service, *MockRepo) {
	repo := new(MockRepo)
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	return NewOrderService(repo, logger), repo
}

func (m *MockRepo) CreateOrder(userID uint) (*entity.Order, error) {
	args := m.Called(userID)
	return args.Get(0).(*entity.Order), args.Error(1)
}

func (m *MockRepo) ChangeOrderStatus(orderID uint, status entity.OrderStatus) (*entity.Order, error) {
	args := m.Called(orderID, status)
	return args.Get(0).(*entity.Order), args.Error(1)
}

func (m *MockRepo) GetUserOrders(userID uint, status string, page, limit int) (*entity.OrderPaginationResponse, error) {
	args := m.Called(userID, status, page, limit)
	return args.Get(0).(*entity.OrderPaginationResponse), args.Error(1)
}

func (m *MockRepo) GetUserOrder(orderID uint) (*entity.Order, error) {
	args := m.Called(orderID)
	return args.Get(0).(*entity.Order), args.Error(1)
}
