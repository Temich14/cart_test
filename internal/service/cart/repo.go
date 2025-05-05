package cart

import "github.com/Temich14/cart_test/internal/domain/entity"

type Repository interface {
	SaveCartItem(item *entity.CartItem) error
	AddProduct(cartID, productID uint, quantity int) (*entity.CartItem, error)
	GetCartID(userID uint) (uint, error)
	GetUserCart(userID uint, page, limit int) (*entity.CartWithItemsPagination, error)
	SaveCart(cart *entity.Cart) error
	RemoveProduct(cartID, productID uint) (uint, error)
	ChangeQuantity(cartID, productID uint, addQuantity int) error
	UpdateTotalQuantity(cartID uint) error
}
