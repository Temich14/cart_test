package repository

import (
	"errors"
	"github.com/Temich14/cart_test/internal/domain/entity"
	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}
func (r *CartRepository) SaveCart(cart *entity.Cart) error {
	return r.db.Save(cart).Error
}
func (r *CartRepository) AddProduct(cartID, productID uint, quantity int) error {
	item := entity.CartItem{CartID: cartID, ProductID: productID, Quantity: quantity}
	if err := r.db.Create(&item).Error; err != nil {
		return err
	}
	return nil
}
func (r *CartRepository) GetUserCart(userID uint) (*entity.Cart, error) {
	var cart entity.Cart
	err := r.db.Where("user_id = ?", userID).First(&cart).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		cart = entity.Cart{
			UserID: userID,
		}
		if err := r.db.Create(&cart).Error; err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}
	return &cart, nil
}
func (r *CartRepository) RemoveProduct(cartID, productID uint) error {
	item := entity.CartItem{CartID: cartID, ProductID: productID}
	if err := r.db.Where(&item).Delete(&item).Error; err != nil {
		return err
	}
	return nil
}
func (r *CartRepository) ChangeQuantity(cartID uint, carItemID uint, newQuantity int) error {
	item := entity.CartItem{CartID: cartID, ProductID: carItemID}
	if err := r.db.Model(&item).UpdateColumn("quantity", newQuantity).Error; err != nil {
		return err
	}
	return nil
}
