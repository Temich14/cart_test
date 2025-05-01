package cart

import "github.com/Temich14/cart_test/internal/domain/entity"

type Service struct {
	repo Repository
}

func NewCartService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AddProductToCart(userID, productID uint, quantity int) error {
	cart, err := s.repo.GetUserCart(userID)
	if err != nil {
		return err
	}
	err = s.repo.AddProduct(cart.ID, productID, quantity)
	if err != nil {
		return err
	}
	err = s.calculateTotalQuantity(cart)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) RemoveProductFromCart(userID, productID uint) error {
	cart, err := s.repo.GetUserCart(userID)
	if err != nil {
		return err
	}
	err = s.repo.RemoveProduct(cart.ID, productID)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) GetUserCart(userID uint) (*entity.Cart, error) {
	cart, err := s.repo.GetUserCart(userID)
	if err != nil {
		return nil, err
	}
	return cart, nil
}
func (s *Service) ChangeQuantity(userID, productID uint, quantity int) error {
	cart, err := s.repo.GetUserCart(userID)
	if err != nil {
		return err
	}
	err = s.repo.ChangeQuantity(cart.ID, productID, quantity)
	if err != nil {
		return err
	}

	err = s.calculateTotalQuantity(cart)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) calculateTotalQuantity(cart *entity.Cart) error {
	totalQuantity := 0
	for _, item := range cart.Items {
		totalQuantity += item.Quantity
	}
	cart.TotalQuantity = totalQuantity
	err := s.repo.SaveCart(cart)
	if err != nil {
		return err
	}
	return nil
}
