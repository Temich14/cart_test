package service

import "github.com/Temich14/cart_test/internal/domain/entity"

// ProductProvider интерфейс для адаптера получения информации о товарах
type ProductProvider interface {
	GetProductByID(productID uint) (*entity.Product, error)
	GetProductsByIDs(productIDs []uint) (map[uint]*entity.Product, error)
}
