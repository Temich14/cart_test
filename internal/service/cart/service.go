package cart

import (
	"github.com/Temich14/cart_test/internal/domain/entity"
)

type Service struct {
	repo Repository
}

func NewCartService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AddProductToCart(userID, productID uint, quantity int) (*entity.CartItem, error) {
	if quantity < 1 {
		return nil, ErrQuantityLessThanZero
	}
	cartID, err := s.repo.GetCartID(userID)
	if err != nil {
		return nil, err
	}
	newItem, err := s.repo.AddProduct(cartID, productID, quantity)
	if err != nil {
		return nil, err
	}
	err = s.repo.UpdateTotalQuantity(cartID)
	if err != nil {
		return nil, err
	}
	return newItem, nil
}
func (s *Service) RemoveProductFromCart(userID, productID uint) (uint, error) {
	cartID, err := s.repo.GetCartID(userID)
	if err != nil {
		return 0, err
	}
	rId, err := s.repo.RemoveProduct(cartID, productID)
	if err != nil {
		return 0, err
	}
	err = s.repo.UpdateTotalQuantity(cartID)
	if err != nil {
		return 0, err
	}
	return rId, nil
}
func (s *Service) GetUserCart(userID uint, page, limit int) (*entity.CartWithItemsPagination, error) {
	cart, err := s.repo.GetUserCart(userID, page, limit)
	if err != nil {
		return nil, err
	}
	return cart, nil
}
func (s *Service) ChangeQuantity(userID, productID uint, quantity int) (int, error) {
	cartID, err := s.repo.GetCartID(userID)
	if err != nil {
		return 0, err
	}
	if quantity <= 0 {
		if _, err = s.RemoveProductFromCart(userID, productID); err != nil {
			return 0, err
		}
		err = s.repo.UpdateTotalQuantity(cartID)
		if err != nil {
			return 0, err
		}
		return 0, nil
	} else {
		if err = s.repo.ChangeQuantity(cartID, productID, quantity); err != nil {
			return 0, err
		}
	}
	err = s.repo.UpdateTotalQuantity(cartID)
	if err != nil {
		return 0, err
	}
	return quantity, nil
}
