// Package repository содержит реализацию взаимодействия с базой данных
// для управления корзинами, товарами и заказами.
package repository

import (
	"errors"
	"fmt"
	"github.com/Temich14/cart_test/internal/config"
	"github.com/Temich14/cart_test/internal/domain/entity"
	glog "github.com/Temich14/cart_test/internal/logger"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"log"
	"log/slog"
	"math"
	"runtime/debug"
	"time"
)

type Repository struct {
	db  *gorm.DB
	cfg *config.DBConfig
}

// CloseDB закрывает соединение с базой данных.
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

// NewRepository создает и возвращает новый экземпляр Repository, устанавливая подключение к БД.
func NewRepository(cfg *config.DBConfig, logger *slog.Logger, env string) *Repository {
	logger.Info("opening db connection")
	gormLog := glog.NewGormLogger(logger, gormlogger.Info)
	if env == "PROD" {
		gormLog = glog.NewGormLogger(logger, gormlogger.Warn)
	}
	db, err := gorm.Open(postgres.Open(cfg.Conn), &gorm.Config{
		Logger: gormLog,
	})
	if err != nil {
		logger.Error("error opening db connection", slog.String("error", err.Error()), slog.String("stack", string(debug.Stack())))
		log.Fatal(err)
	}
	logger.Info("db connection opened")

	return &Repository{
		db: db, cfg: cfg,
	}
}

// SaveCartItem сохраняет или обновляет элемент корзины.
func (r *Repository) SaveCartItem(item *entity.CartItem) error {
	return r.db.Save(item).Error
}

// SaveCart сохраняет или обновляет корзину.
func (r *Repository) SaveCart(cart *entity.Cart) error {
	return r.db.Save(cart).Error
}

// AddProduct добавляет товар в корзину или увеличивает его количество, если товар уже есть.
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

// GetCartID получает ID корзины пользователя, создавая ее при необходимости.
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

// UpdateTotalQuantity обновляет поле total_quantity корзины на основе суммы всех товаров.
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

// GetCartMeta возвращает метаинформацию о корзине без списка товаров.
func (r *Repository) GetCartMeta(cartID uint) (*entity.Cart, error) {
	cart := entity.Cart{
		ID: cartID,
	}
	err := r.db.Model(&entity.Cart{}).
		Select("id", "user_id", "total_quantity", "total_cost", "created_at", "updated_at").
		Where("id = ?", cartID).
		First(&cart).Error
	return &cart, err
}

// GetUserCart возвращает корзину пользователя с пагинацией товаров.
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
		TotalCost:     cart.TotalCost,
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

// RemoveProduct удаляет товар из корзины.
func (r *Repository) RemoveProduct(cartID, productID uint) (*entity.CartItem, error) {
	var item entity.CartItem
	err := r.db.Where("cart_id = ? AND product_id = ?", cartID, productID).First(&item).Error
	if err != nil {
		return nil, err
	}
	err = r.db.Where("cart_id = ? AND product_id = ?", cartID, productID).Delete(&entity.CartItem{}).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// ChangeQuantity изменяет количество определенного товара в корзине.
func (r *Repository) ChangeQuantity(cartID uint, carItemID uint, newQuantity int) (*entity.CartItem, error) {
	prevItem := entity.CartItem{
		ID: carItemID,
	}
	err := r.db.Model(&prevItem).Where("id = ?", carItemID).First(&prevItem).Error
	if err != nil {
		return nil, err
	}
	if err = r.db.Model(&entity.CartItem{CartID: cartID, ProductID: carItemID}).
		Where("cart_id = ? AND product_id = ?", cartID, carItemID).
		UpdateColumn("quantity", newQuantity).Error; err != nil {
		return nil, err
	}
	return &prevItem, nil
}

// CreateOrder создает заказ из корзины пользователя.
func (r *Repository) CreateOrder(userID uint) (*entity.Order, error) {
	var cart entity.Cart
	if err := r.db.Preload("Items").Where("user_id = ?", userID).First(&cart).Error; err != nil {
		return nil, err
	}
	status := entity.CREATED
	order := entity.Order{
		UserID:        cart.UserID,
		Cost:          cart.TotalCost,
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

// ChangeOrderStatus изменяет статус существующего заказа.
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

// GetUserOrders возвращает заказы пользователя с пагинацией и фильтрацией по статусу.
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

// GetUserOrder возвращает конкретный заказ по ID, включая его товары.
func (r *Repository) GetUserOrder(orderID uint) (*entity.Order, error) {
	var order *entity.Order
	if err := r.db.Preload("Items").Where("id = ?", orderID).First(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

// GetProductsByIDs возвращает мапу продуктов по их ID.
func (r *Repository) GetProductsByIDs(productIDs []uint) (map[uint]*entity.Product, error) {
	if len(productIDs) == 0 {
		return map[uint]*entity.Product{}, nil
	}

	var products []entity.Product
	err := r.db.Where("id IN ?", productIDs).Find(&products).Error
	if err != nil {
		return nil, err
	}

	productMap := make(map[uint]*entity.Product, len(products))
	for i := range products {
		p := products[i]
		productMap[p.ID] = &p
	}
	return productMap, nil
}

// GetProductByID возвращает продукт по его ID.
func (r *Repository) GetProductByID(productID uint) (*entity.Product, error) {
	var product entity.Product
	err := r.db.Where("id = ?", productID).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("product with ID %d not found", productID)
		}
		return nil, err
	}
	return &product, nil
}

// paginate отвечает за пагинацию
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
