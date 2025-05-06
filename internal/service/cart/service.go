package cart

import (
	"github.com/Temich14/cart_test/internal/domain/entity"
	"log/slog"
	"runtime/debug"
)

type Service struct {
	repo Repository
	log  *slog.Logger
}

func NewCartService(repo Repository, log *slog.Logger) *Service {
	return &Service{repo: repo, log: log}
}

func (s *Service) AddProductToCart(userID, productID uint, quantity int) (*entity.CartItem, error) {
	s.log.Debug(
		"adding product to cart",
		slog.Uint64("user_id", uint64(userID)),
		slog.Uint64("product_id", uint64(productID)),
		slog.Int("quantity", quantity))
	if quantity < 1 {
		s.log.Error(
			"quantity must be > 0",
			slog.Uint64("user_id", uint64(userID)),
			slog.Uint64("product_id", uint64(productID)))
		return nil, ErrQuantityLessThanZero
	}

	cartID, err := s.repo.GetCartID(userID)
	if err != nil {
		s.log.Error(
			"error getting cart",
			slog.Uint64("user_id", uint64(userID)),
			slog.Uint64("product_id", uint64(productID)),
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return nil, err
	}
	newItem, err := s.repo.AddProduct(cartID, productID, quantity)
	if err != nil {
		s.log.Error(
			"error adding product to cart",
			slog.Uint64("user_id", uint64(userID)),
			slog.Uint64("product_id", uint64(productID)),
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return nil, err
	}
	err = s.repo.UpdateTotalQuantity(cartID)
	if err != nil {
		s.log.Error(
			"error updating total quantity",
			slog.Uint64("user_id", uint64(userID)),
			slog.Uint64("product_id", uint64(productID)),
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return nil, err
	}
	s.log.Debug("item added to cart", slog.Uint64("user_id", uint64(userID)), slog.Uint64("product_id", uint64(productID)))
	return newItem, nil
}
func (s *Service) RemoveProductFromCart(userID, productID uint) (uint, error) {
	s.log.Debug("removing product from cart",
		slog.Uint64("user_id", uint64(userID)),
		slog.Uint64("product_id", uint64(productID)))

	cartID, err := s.repo.GetCartID(userID)
	if err != nil {
		s.log.Error(
			"failed to get cart ID",
			slog.Uint64("user_id", uint64(userID)),
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return 0, err
	}
	rId, err := s.repo.RemoveProduct(cartID, productID)
	if err != nil {
		s.log.Error(
			"failed to remove product",
			slog.Uint64("cart_id", uint64(cartID)),
			slog.Uint64("product_id", uint64(productID)),
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return 0, err
	}
	err = s.repo.UpdateTotalQuantity(cartID)
	if err != nil {
		s.log.Error(
			"failed to update total quantity",
			slog.Uint64("cart_id", uint64(cartID)),
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return 0, err
	}

	s.log.Debug("product removed from cart", slog.Uint64("user_id", uint64(userID)), slog.Uint64("product_id", uint64(productID)))
	return rId, nil
}
func (s *Service) GetUserCart(userID uint, page, limit int) (*entity.CartWithItemsPagination, error) {
	s.log.Debug("retrieving user cart",
		slog.Uint64("user_id", uint64(userID)),
		slog.Int("page", page),
		slog.Int("limit", limit))

	cart, err := s.repo.GetUserCart(userID, page, limit)
	if err != nil {
		s.log.Error(
			"failed to retrieve cart",
			slog.Uint64("user_id", uint64(userID)),
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return nil, err
	}

	s.log.Debug("user cart retrieved", slog.Uint64("user_id", uint64(userID)))
	return cart, nil
}
func (s *Service) ChangeQuantity(userID, productID uint, quantity int) (int, error) {
	s.log.Debug("changing product quantity",
		slog.Uint64("user_id", uint64(userID)),
		slog.Uint64("product_id", uint64(productID)),
		slog.Int("quantity", quantity))

	cartID, err := s.repo.GetCartID(userID)
	if err != nil {
		s.log.Error(
			"failed to get cart ID",
			slog.Uint64("user_id", uint64(userID)),
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return 0, err
	}
	if quantity <= 0 {
		s.log.Debug("quantity <= 0, removing product", slog.Uint64("user_id", uint64(userID)), slog.Uint64("product_id", uint64(productID)))
		if _, err = s.RemoveProductFromCart(userID, productID); err != nil {
			s.log.Error(
				"failed to remove product",
				slog.Uint64("user_id", uint64(userID)),
				slog.String("error", err.Error()),
				slog.String("stack", string(debug.Stack())))
			return 0, err
		}
		err = s.repo.UpdateTotalQuantity(cartID)
		if err != nil {
			s.log.Error(
				"failed to update total quantity after removal",
				slog.Uint64("cart_id", uint64(cartID)),
				slog.String("error", err.Error()),
				slog.String("stack", string(debug.Stack())))
			return 0, err
		}
		s.log.Debug("product removed due to zero quantity", slog.Uint64("user_id", uint64(userID)), slog.Uint64("product_id", uint64(productID)))
		return 0, nil
	}
	if err = s.repo.ChangeQuantity(cartID, productID, quantity); err != nil {
		s.log.Error(
			"failed to change product quantity",
			slog.Uint64("cart_id", uint64(cartID)),
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return 0, err
	}
	err = s.repo.UpdateTotalQuantity(cartID)
	if err != nil {
		s.log.Error(
			"failed to update total quantity after change",
			slog.Uint64("cart_id", uint64(cartID)),
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return 0, err
	}
	s.log.Debug("product quantity changed", slog.Uint64("user_id", uint64(userID)), slog.Uint64("product_id", uint64(productID)), slog.Int("quantity", quantity))
	return quantity, nil
}
