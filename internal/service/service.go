package service

import "github.com/Temich14/cart_test/internal/domain/entity"

type CartRepository interface {
	AddProduct(cartID, productID uint, quantity int) error
	GetUserCart(userID uint) (*entity.Cart, error)
	SaveCart(cart *entity.Cart) error
}
type CartService struct {
	repo CartRepository
}

func NewCartService(repo CartRepository) *CartService {
	return &CartService{repo: repo}
}

func (s *CartService) AddProductToCart(userID, productID uint, quantity int) error {
	cart, err := s.repo.GetUserCart(userID)
	if err != nil {
		return err
	}
	err = s.repo.AddProduct(cart.ID, productID, quantity)
	if err != nil {
		return err
	}
	cart.TotalQuantity += quantity
	err = s.repo.SaveCart(cart)
	if err != nil {
		return err
	}
	return nil
}
