package service

import "github.com/Temich14/cart_test/internal/domain/entity"

type CartRepository interface {
	AddProduct(cartID, productID uint, quantity int) error
	GetUserCart(userID uint) (*entity.Cart, error)
	SaveCart(cart *entity.Cart) error
	RemoveProduct(cartID, productID uint) error
	ChangeQuantity(cartID, productID uint, addQuantity int) error
}
