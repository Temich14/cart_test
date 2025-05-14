package cart

import (
	"github.com/Temich14/cart_test/internal/domain/entity"
	"github.com/stretchr/testify/mock"
	"io"
	"log/slog"
)

type MockRepository struct {
	mock.Mock
}

func newServiceWithMock() (*Service, *MockRepository) {
	mockRepo := new(MockRepository)
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	service := NewCartService(mockRepo, logger, mockRepo)
	return service, mockRepo
}

func (m *MockRepository) SaveCartItem(item *entity.CartItem) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *MockRepository) AddProduct(cartID, productID uint, quantity int) (*entity.CartItem, error) {
	args := m.Called(cartID, productID, quantity)
	return args.Get(0).(*entity.CartItem), args.Error(1)
}

func (m *MockRepository) GetCartID(userID uint) (uint, error) {
	args := m.Called(userID)
	return uint(args.Int(0)), args.Error(1)
}

func (m *MockRepository) GetUserCart(userID uint, page, limit int) (*entity.CartWithItemsPagination, error) {
	args := m.Called(userID, page, limit)
	return args.Get(0).(*entity.CartWithItemsPagination), args.Error(1)
}

func (m *MockRepository) SaveCart(cart *entity.Cart) error {
	args := m.Called(cart)
	return args.Error(0)
}

func (m *MockRepository) RemoveProduct(cartID, productID uint) (*entity.CartItem, error) {
	args := m.Called(cartID, productID)
	return args.Get(0).(*entity.CartItem), args.Error(1)
}

func (m *MockRepository) ChangeQuantity(cartID, productID uint, quantity int) (*entity.CartItem, error) {
	args := m.Called(cartID, productID, quantity)
	return args.Get(0).(*entity.CartItem), args.Error(0)
}

func (m *MockRepository) UpdateTotalQuantity(cartID uint) error {
	args := m.Called(cartID)
	return args.Error(0)
}
func (m *MockRepository) UpdateTotalCost(cartID uint, newCost float32) error {
	args := m.Called(cartID)
	return args.Error(0)
}
func (m *MockRepository) GetProductByID(productID uint) (*entity.Product, error) {
	return nil, nil
}
func (m *MockRepository) GetProductsByIDs(productIDs []uint) (map[uint]*entity.Product, error) {
	return nil, nil
}
func (m *MockRepository) GetCartMeta(cartID uint) (*entity.Cart, error) {
	args := m.Called(cartID)
	return args.Get(0).(*entity.Cart), args.Error(1)
}
