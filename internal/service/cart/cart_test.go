package cart

import (
	"errors"
	"github.com/Temich14/cart_test/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/mock"
	"testing"
)

var errAny = errors.New("some error")

func TestAddProductToCart(t *testing.T) {
	userID := uint(1)
	productID := uint(2)
	qty := 3
	cartID := uint(5)
	item := &entity.CartItem{ID: 1, ProductID: productID, Quantity: qty}

	t.Run("success", func(t *testing.T) {
		service, mockRepo := newServiceWithMock()
		mockRepo.On("GetCartID", userID).Return(int(cartID), nil)
		mockRepo.On("AddProduct", cartID, productID, qty).Return(item, nil)
		mockRepo.On("UpdateTotalQuantity", cartID).Return(nil)

		result, err := service.AddProductToCart(userID, productID, qty)
		assert.NoError(t, err)
		assert.Equal(t, item, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("quantity zero", func(t *testing.T) {
		service, mockRepo := newServiceWithMock()
		mockRepo.On("GetCartID", userID).Return(int(cartID), nil)
		mockRepo.On("AddProduct", cartID, productID, qty).Return(item, nil)
		mockRepo.On("UpdateTotalQuantity", cartID).Return(nil)
		result, err := service.AddProductToCart(userID, productID, 0)
		assert.ErrorIs(t, err, ErrQuantityLessThanZero)
		assert.Nil(t, result)
	})

	t.Run("AddProduct error", func(t *testing.T) {
		service, mockRepo := newServiceWithMock()
		mockRepo.On("GetCartID", userID).Return(int(cartID), nil)
		mockRepo.On("AddProduct", cartID, productID, qty).Return((*entity.CartItem)(nil), errAny)
		mockRepo.On("UpdateTotalQuantity", cartID).Return(nil)
		result, err := service.AddProductToCart(userID, productID, qty)
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("UpdateTotalQuantity error", func(t *testing.T) {
		service, mockRepo := newServiceWithMock()
		mockRepo.On("GetCartID", userID).Return(int(cartID), nil)
		mockRepo.On("AddProduct", cartID, productID, qty).Return(item, nil)
		mockRepo.On("UpdateTotalQuantity", cartID).Return(errAny)

		result, err := service.AddProductToCart(userID, productID, qty)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}
func TestRemoveProductFromCart(t *testing.T) {
	userID := uint(1)
	productID := uint(2)
	cartID := uint(3)

	t.Run("success", func(t *testing.T) {
		service, mockRepo := newServiceWithMock()
		mockRepo.On("GetCartID", userID).Return(int(cartID), nil)
		mockRepo.On("RemoveProduct", cartID, productID).Return(int(productID), nil)
		mockRepo.On("UpdateTotalQuantity", cartID).Return(nil)

		id, err := service.RemoveProductFromCart(userID, productID)
		assert.NoError(t, err)
		assert.Equal(t, productID, id)
	})
}

func TestChangeQuantity(t *testing.T) {
	userID := uint(1)
	productID := uint(2)
	cartID := uint(3)

	t.Run("success", func(t *testing.T) {
		service, mockRepo := newServiceWithMock()
		mockRepo.On("GetCartID", userID).Return(int(cartID), nil)
		mockRepo.On("ChangeQuantity", cartID, productID, 5).Return(nil)
		mockRepo.On("RemoveProduct", cartID, productID).Return(int(productID), nil)
		mockRepo.On("UpdateTotalQuantity", cartID).Return(nil)

		qty, err := service.ChangeQuantity(userID, productID, 5)
		assert.NoError(t, err)
		assert.Equal(t, 5, qty)
	})

	t.Run("zero quantity removes item", func(t *testing.T) {
		service, mockRepo := newServiceWithMock()
		mockRepo.On("GetCartID", userID).Return(int(cartID), nil)
		mockRepo.On("ChangeQuantity", cartID, productID, 5).Return(nil)
		mockRepo.On("RemoveProduct", cartID, productID).Return(int(productID), nil)
		mockRepo.On("UpdateTotalQuantity", cartID).Return(nil)

		qty, err := service.ChangeQuantity(userID, productID, 0)
		assert.NoError(t, err)
		assert.Equal(t, 0, qty)
	})
}

func TestGetUserCart(t *testing.T) {
	userID := uint(1)
	page := 1
	limit := 10
	expected := &entity.CartWithItemsPagination{TotalQuantity: 1}

	t.Run("success", func(t *testing.T) {
		service, mockRepo := newServiceWithMock()
		mockRepo.On("GetUserCart", userID, page, limit).Return(expected, nil)

		result, err := service.GetUserCart(userID, page, limit)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})
}
