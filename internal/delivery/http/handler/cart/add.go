package cart

import (
	"errors"
	"github.com/Temich14/cart_test/internal/delivery/http/handler/utils"
	"github.com/Temich14/cart_test/internal/service/cart"
	"github.com/gin-gonic/gin"
	"net/http"
)

type addDTO struct {
	ProductID uint `json:"product_id" example:"1"`
	Quantity  int  `json:"quantity" example:"2"`
}

// Add godoc
//
//	@Summary		Добавить товар в корзину
//	@Description	Добавляет товар с указанным id и количеством в корзину пользователя и возвращает обновленную корзину. В случае, если в корзине уже есть товар с данным id - увеличит количество товара в корзине на указанное число.
//	@Tags			cart
//	@Accept			json
//	@Produce		json
//
//	@Param			user_id	query		uint	false	"id пользователя"	example(1)
//
//	@Param			input	body		addDTO	true	"Данные о товаре"
//	@Success		201		{object}	entity.Cart
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/cart [post]
func (h *Handler) Add(c *gin.Context) {
	var dto addDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, err := utils.TryGetUserID(c)
	if err != nil {
		c.Abort()
		return
	}
	newCart, err := h.s.AddProductToCart(userID, dto.ProductID, dto.Quantity)
	if err != nil {
		if errors.Is(err, cart.ErrQuantityLessThanZero) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"cart": *newCart})
}
