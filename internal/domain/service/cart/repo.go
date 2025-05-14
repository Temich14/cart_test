package cart

import "github.com/Temich14/cart_test/internal/domain/entity"

// Repository интерфейс для взаимодействия с бд.
type Repository interface {
	SaveCartItem(item *entity.CartItem) error
	AddProduct(cartID, productID uint, quantity int) (*entity.CartItem, error)
	GetCartID(userID uint) (uint, error)
	GetUserCart(userID uint, page, limit int) (*entity.CartWithItemsPagination, error)
	SaveCart(cart *entity.Cart) error
	RemoveProduct(cartID, productID uint) (*entity.CartItem, error)
	ChangeQuantity(cartID, productID uint, addQuantity int) (*entity.CartItem, error)
	GetCartMeta(cartID uint) (*entity.Cart, error)
}
