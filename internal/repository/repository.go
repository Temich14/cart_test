package repository

import (
	"errors"
	"github.com/Temich14/cart_test/internal/config"
	"github.com/Temich14/cart_test/internal/domain/entity"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"log/slog"
	"time"
)

type Repository struct {
	db     *gorm.DB
	cfg    *config.DBConfig
	logger *slog.Logger
}

func (r *Repository) CloseDB() error {
	db, err := r.db.DB()
	if err != nil {
		return err
	}
	if err = db.Close(); err != nil {
		return err
	}
	return nil
}
func NewRepository(cfg *config.DBConfig, logger *slog.Logger) *Repository {
	db, err := gorm.Open(postgres.Open(cfg.Conn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return &Repository{
		db: db, cfg: cfg, logger: logger,
	}
}
func (r *Repository) SaveCart(cart *entity.Cart) error {
	return r.db.Save(cart).Error
}
func (r *Repository) AddProduct(cartID, productID uint, quantity int) error {
	item := entity.CartItem{CartID: cartID, ProductID: productID, Quantity: quantity}
	if err := r.db.Create(&item).Error; err != nil {
		return err
	}
	return nil
}
func (r *Repository) GetUserCart(userID uint) (*entity.Cart, error) {
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
func (r *Repository) RemoveProduct(cartID, productID uint) error {
	item := entity.CartItem{CartID: cartID, ProductID: productID}
	if err := r.db.Where(&item).Delete(&item).Error; err != nil {
		return err
	}
	return nil
}
func (r *Repository) ChangeQuantity(cartID uint, carItemID uint, newQuantity int) error {
	item := entity.CartItem{CartID: cartID, ProductID: carItemID}
	if err := r.db.Model(&item).UpdateColumn("quantity", newQuantity).Error; err != nil {
		return err
	}
	return nil
}
func (r *Repository) CreateOrder(cart *entity.Cart) (*entity.Order, error) {
	status := entity.CREATED
	order := entity.Order{
		UserID:        cart.UserID,
		ItemsQuantity: cart.TotalQuantity,
		Status:        string(status),
		CreatedAt:     time.Now(),
	}
	if err := r.db.Create(&entity.Order{}).Error; err != nil {
		return nil, err
	}
	for _, cartItem := range cart.Items {
		orderItem := entity.OrderItem{
			ProductID: cartItem.ProductID,
			Quantity:  cartItem.Quantity,
		}
		order.Items = append(order.Items, orderItem)
	}
	if err := r.db.Save(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
func (r *Repository) ChangeOrderStatus(orderID uint, status entity.OrderStatus) error {
	var order *entity.Order
	err := r.db.Where("order_id = ?", orderID).First(&order).Error
	if err != nil {
		return err
	}
	return r.db.Model(order).Update("status", string(status)).Error
}
func (r *Repository) GetUserOrders(userID uint) ([]*entity.Order, error) {
	var orders []*entity.Order
	err := r.db.Where("user_id = ?", userID).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
func (r *Repository) GetUserOrder(orderID uint) (*entity.Order, error) {
	var order *entity.Order
	if err := r.db.Where("order_id = ?", orderID).First(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}
