// Package cart реализует бизнес-логику для управления корзиной пользователя.
package cart

import (
	"github.com/Temich14/cart_test/internal/domain/entity"
	"github.com/Temich14/cart_test/internal/domain/service"
	"log/slog"
	"runtime/debug"
)

type Service struct {
	repo            Repository
	log             *slog.Logger
	productProvider service.ProductProvider
}

// NewCartService создает новый экземпляр сервиса корзины.
//   - repo: интерфейс репозитория для доступа к данным корзины.
//   - log: логгер.
//   - provider: интерфейс адаптера для получения информации о продуктах.
func NewCartService(repo Repository, log *slog.Logger, provider service.ProductProvider) *Service {
	return &Service{repo: repo, log: log, productProvider: provider}
}

// AddProductToCart добавляет товар в корзину пользователя.
// Если товар уже в корзине — увеличивает его количество.
// Возвращает добавленный элемент корзины или ошибку.
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

	product, err := s.productProvider.GetProductByID(productID)
	if err != nil {
		s.log.Error(
			"error getting product",
			slog.Uint64("user_id", uint64(userID)),
			slog.Uint64("product_id", uint64(productID)),
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return nil, err
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
	cartMeta, err := s.repo.GetCartMeta(cartID)
	if err != nil {
		s.log.Error(
			"error getting cartMeta",
			slog.Uint64("user_id", uint64(userID)),
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return nil, err
	}
	cartMeta.TotalQuantity += quantity
	cartMeta.TotalCost += product.Cost * float32(quantity)
	err = s.repo.SaveCart(cartMeta)
	if err != nil {
		s.log.Error(
			"error saving cart",
			slog.Uint64("user_id", uint64(userID)),
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return nil, err
	}
	s.log.Debug("item added to cart", slog.Uint64("user_id", uint64(userID)), slog.Uint64("product_id", uint64(productID)))
	return newItem, nil
}

// RemoveProductFromCart удаляет товар из корзины пользователя.
// Возвращает ID удалённого элемента или ошибку.
func (s *Service) RemoveProductFromCart(userID, productID uint) (uint, error) {
	s.log.Debug("removing product from cart",
		slog.Uint64("user_id", uint64(userID)),
		slog.Uint64("product_id", uint64(productID)))

	product, err := s.productProvider.GetProductByID(productID)
	if err != nil {
		s.log.Error(
			"error getting product",
			slog.Uint64("user_id", uint64(userID)),
			slog.Uint64("product_id", uint64(productID)),
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return 0, err
	}
	cartID, err := s.repo.GetCartID(userID)
	if err != nil {
		s.log.Error(
			"failed to get cart ID",
			slog.Uint64("user_id", uint64(userID)),
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return 0, err
	}
	cartItem, err := s.repo.RemoveProduct(cartID, productID)
	if err != nil {
		s.log.Error(
			"failed to remove product",
			slog.Uint64("cart_id", uint64(cartID)),
			slog.Uint64("product_id", uint64(productID)),
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return 0, err
	}
	cartMeta, err := s.repo.GetCartMeta(cartID)
	if err != nil {
		s.log.Error(
			"error getting cartMeta",
			slog.Uint64("user_id", uint64(userID)),
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return 0, err
	}
	cartMeta.TotalQuantity -= cartItem.Quantity
	cartMeta.TotalCost -= product.Cost * float32(cartItem.Quantity)
	err = s.repo.SaveCart(cartMeta)
	if err != nil {
		s.log.Error(
			"error saving cart",
			slog.Uint64("user_id", uint64(userID)),
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return 0, err
	}

	s.log.Debug("product removed from cart", slog.Uint64("user_id", uint64(userID)), slog.Uint64("product_id", uint64(productID)))
	return cartItem.ID, nil
}

// GetUserCart возвращает корзину пользователя с пагинацией и полной информацией о товарах.
// Возвращает структуру с товарами и метаинформацией или ошибку.
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

	var productsIDs []uint

	for _, item := range cart.Items {
		productsIDs = append(productsIDs, item.ProductID)
	}

	itemsMap, err := s.productProvider.GetProductsByIDs(productsIDs)

	if err != nil {
		return nil, err
	}
	for i := range cart.Items {
		item := &cart.Items[i]
		product, ok := itemsMap[item.ProductID]
		if !ok {
			s.log.Warn("product getting error", slog.Uint64("product_id", uint64(item.ProductID)))
			continue
		}
		item.Product = *product
	}
	s.log.Debug("user cart retrieved", slog.Uint64("user_id", uint64(userID)))
	return cart, nil
}

// ChangeQuantity изменяет количество товара в корзине пользователя.
// Если количество становится <= 0 — товар удаляется из корзины.
// Возвращает новое количество или ошибку.
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
	product, err := s.productProvider.GetProductByID(productID)
	if err != nil {
		return 0, err
	}
	cartMeta, err := s.repo.GetCartMeta(cartID)
	if err != nil {
		return 0, err
	}
	if quantity <= 0 {
		s.log.Debug("quantity <= 0, removing product", slog.Uint64("user_id", uint64(userID)), slog.Uint64("product_id", uint64(productID)))
		_, err = s.RemoveProductFromCart(userID, productID)
		if err != nil {
			s.log.Error(
				"failed to remove product",
				slog.Uint64("user_id", uint64(userID)),
				slog.String("error", err.Error()),
				slog.String("stack", string(debug.Stack())))
			return 0, err
		}
		s.log.Debug("product removed due to zero quantity", slog.Uint64("user_id", uint64(userID)), slog.Uint64("product_id", uint64(productID)))
		return 0, nil
	}
	prevItem, err := s.repo.ChangeQuantity(cartID, productID, quantity)
	if err != nil {
		s.log.Error(
			"failed to change product quantity",
			slog.Uint64("cart_id", uint64(cartID)),
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return 0, err
	}
	delta := quantity - prevItem.Quantity
	cartMeta.TotalQuantity += delta
	cartMeta.TotalCost += product.Cost * float32(delta)

	s.log.Info("delta", "delta", delta, "cartMeta", cartMeta)
	s.log.Info("delta", "prev", prevItem.Quantity, "quanitty", quantity)

	err = s.repo.SaveCart(cartMeta)
	if err != nil {
		s.log.Error(
			"error saving cart",
			slog.Uint64("cart_id", uint64(cartID)),
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return 0, err
	}

	s.log.Debug("product quantity changed", slog.Uint64("user_id", uint64(userID)), slog.Uint64("product_id", uint64(productID)), slog.Int("quantity", quantity))
	return quantity, nil
}
