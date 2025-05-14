package cart

import "github.com/Temich14/cart_test/internal/domain/entity"

type CartService interface {
	AddProductToCart(userID, productID uint, quantity int) (*entity.CartItem, error)
	RemoveProductFromCart(userID, productID uint) (uint, error)
	GetUserCart(userID uint, page, limit int) (*entity.CartWithItemsPagination, error)
	ChangeQuantity(userID, productID uint, quantity int) (int, error)
}
