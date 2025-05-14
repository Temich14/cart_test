package order

import (
	"errors"
	"github.com/Temich14/cart_test/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewOrder_Success(t *testing.T) {
	service, repo := NewServiceWithMock()
	expectedOrder := &entity.Order{ID: 1, UserID: 123}
	repo.On("CreateOrder", uint(123)).Return(expectedOrder, nil)

	order, err := service.CreateNewOrder(123)

	assert.NoError(t, err)
	assert.Equal(t, expectedOrder, order)
	repo.AssertExpectations(t)
}

func TestCreateNewOrder_Error(t *testing.T) {
	service, repo := NewServiceWithMock()

	repo.On("CreateOrder", uint(123)).Return(&entity.Order{}, errors.New("db error"))

	order, err := service.CreateNewOrder(123)

	assert.Error(t, err)
	assert.Nil(t, order)
	repo.AssertExpectations(t)
}
func TestChangeStatus_Success(t *testing.T) {
	service, repo := NewServiceWithMock()

	expectedOrder := &entity.Order{ID: 1, Status: string(entity.CREATED)}
	repo.On("ChangeOrderStatus", uint(1), entity.CREATED).Return(expectedOrder, nil)

	order, err := service.ChangeStatus(1, entity.CREATED)

	assert.NoError(t, err)
	assert.Equal(t, expectedOrder, order)
	repo.AssertExpectations(t)
}

func TestChangeStatus_Error(t *testing.T) {
	service, repo := NewServiceWithMock()

	repo.On("ChangeOrderStatus", uint(1), entity.CREATED).Return(&entity.Order{}, errors.New("update failed"))

	order, err := service.ChangeStatus(1, entity.CREATED)

	assert.Error(t, err)
	assert.Nil(t, order)
	repo.AssertExpectations(t)
}
func TestGetOrders_Success(t *testing.T) {
	service, repo := NewServiceWithMock()

	expectedResp := &entity.OrderPaginationResponse{Data: []*entity.Order{{ID: 1}}, Meta: entity.PaginationMeta{Total: 1}}
	repo.On("GetUserOrders", uint(123), "paid", 1, 10).Return(expectedResp, nil)

	resp, err := service.GetOrders(123, "paid", 1, 10)

	assert.NoError(t, err)
	assert.Equal(t, expectedResp, resp)
	repo.AssertExpectations(t)
}

func TestGetOrders_Error(t *testing.T) {
	service, repo := NewServiceWithMock()

	repo.On("GetUserOrders", uint(123), "paid", 1, 10).Return(&entity.OrderPaginationResponse{}, errors.New("query failed"))

	resp, err := service.GetOrders(123, "paid", 1, 10)

	assert.Error(t, err)
	assert.Nil(t, resp)
	repo.AssertExpectations(t)
}
func TestGetOrder_Success(t *testing.T) {
	service, repo := NewServiceWithMock()

	expectedOrder := &entity.Order{ID: 1}
	repo.On("GetUserOrder", uint(1)).Return(expectedOrder, nil)

	order, err := service.GetOrder(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedOrder, order)
	repo.AssertExpectations(t)
}

func TestGetOrder_Error(t *testing.T) {
	service, repo := NewServiceWithMock()

	repo.On("GetUserOrder", uint(1)).Return(&entity.Order{}, errors.New("not found"))

	order, err := service.GetOrder(1)

	assert.Error(t, err)
	assert.Nil(t, order)
	repo.AssertExpectations(t)
}
