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
	"math"
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
func (r *Repository) SaveCartItem(item *entity.CartItem) error {
	return r.db.Save(item).Error
}
func (r *Repository) SaveCart(cart *entity.Cart) error {
	return r.db.Save(cart).Error
}
func (r *Repository) AddProduct(cartID, productID uint, quantity int) (*entity.CartItem, error) {
	var item entity.CartItem
	err := r.db.Where("cart_id = ? AND product_id = ?", cartID, productID).First(&item).Error
	if err == nil {
		item.Quantity += quantity
		if err := r.db.Save(&item).Error; err != nil {
			return nil, err
		}
		return &item, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		newItem := &entity.CartItem{
			CartID:    cartID,
			ProductID: productID,
			Quantity:  quantity,
		}
		if err := r.db.Create(newItem).Error; err != nil {
			return nil, err
		}
		return newItem, nil
	}
	return nil, err
}
func (r *Repository) GetCartID(userID uint) (uint, error) {
	var cart entity.Cart
	err := r.db.Model(&entity.Cart{}).Where("user_id = ?", userID).Select("id").First(&cart).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		cart = entity.Cart{UserID: userID}
		if err := r.db.Create(&cart).Error; err != nil {
			return cart.ID, err
		}
	} else if err != nil {
		return 0, nil
	}
	return cart.ID, nil
}
func (r *Repository) UpdateTotalQuantity(cartID uint) error {
	var totalQuantity int
	err := r.db.Model(&entity.CartItem{}).
		Where("cart_id = ?", cartID).
		Select("SUM(quantity)").
		Scan(&totalQuantity).Error
	if err != nil {
		return err
	}
	return r.db.Model(&entity.Cart{}).
		Where("id = ?", cartID).
		Update("total_quantity", totalQuantity).Error
}
func (r *Repository) GetUserCart(userID uint, page, limit int) (*entity.CartWithItemsPagination, error) {
	var cart entity.Cart
	err := r.db.Where("user_id = ?", userID).First(&cart).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		cart = entity.Cart{UserID: userID}
		if err := r.db.Create(&cart).Error; err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	var total int64
	err = r.db.Model(&entity.CartItem{}).
		Where("cart_id = ?", cart.ID).
		Count(&total).Error
	if err != nil {
		return nil, err
	}
	var items []entity.CartItem
	err = r.db.Scopes(paginate(page, limit)).
		Where("cart_id = ?", cart.ID).
		Find(&items).Error
	if err != nil {
		return nil, err
	}
	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	return &entity.CartWithItemsPagination{
		ID:            cart.ID,
		UserID:        cart.UserID,
		TotalQuantity: cart.TotalQuantity,
		CreatedAt:     cart.CreatedAt,
		UpdatedAt:     cart.UpdatedAt,
		Items:         items,
		ItemsMeta: entity.PaginationMeta{
			Page:       page,
			Limit:      limit,
			Total:      total,
			TotalPages: totalPages,
		},
	}, nil
}
func (r *Repository) RemoveProduct(cartID, productID uint) (uint, error) {
	item := entity.CartItem{CartID: cartID, ProductID: productID}
	if err := r.db.Where(&item).Delete(&item).Error; err != nil {
		return 0, err
	}
	return productID, nil
}
func (r *Repository) ChangeQuantity(cartID uint, carItemID uint, newQuantity int) error {
	if err := r.db.Model(&entity.CartItem{CartID: cartID, ProductID: carItemID}).
		Where("cart_id = ? AND product_id = ?", cartID, carItemID).
		UpdateColumn("quantity", newQuantity).Error; err != nil {
		return err
	}
	return nil
}
func (r *Repository) CreateOrder(userID uint) (*entity.Order, error) {
	var cart entity.Cart
	if err := r.db.Preload("Items").Where("user_id = ?", userID).First(&cart).Error; err != nil {
		return nil, err
	}
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
func (r *Repository) ChangeOrderStatus(orderID uint, status entity.OrderStatus) (*entity.Order, error) {
	if err := r.db.Model(&entity.Order{}).
		Where("id = ?", orderID).
		Update("status", string(status)).Error; err != nil {
		return nil, err
	}
	var updatedOrder *entity.Order
	if err := r.db.Where("id = ?", orderID).First(updatedOrder).Error; err != nil {
		return nil, err
	}
	return updatedOrder, nil
}
func (r *Repository) GetUserOrders(userID uint, status string, page, limit int) (*entity.OrderPaginationResponse, error) {
	var orders []*entity.Order
	var total int64
	db := r.db.Model(&entity.Order{}).Where("user_id = ?", userID)
	if status != "" {
		db = db.Where("status = ?", status)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, err
	}
	if err := db.Scopes(paginate(page, limit)).Preload("Items").Find(&orders).Error; err != nil {
		return nil, err
	}
	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return &entity.OrderPaginationResponse{
		Data: orders,
		Meta: entity.PaginationMeta{
			Total:      total,
			Page:       page,
			Limit:      limit,
			TotalPages: totalPages,
		},
	}, nil
}
func (r *Repository) GetUserOrder(orderID uint) (*entity.Order, error) {
	var order *entity.Order
	if err := r.db.Preload("Items").Where("id = ?", orderID).First(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}
func paginate(page, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 || limit <= 0 {
			return db
		}
		if limit > 100 {
			limit = 100
		}
		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}
