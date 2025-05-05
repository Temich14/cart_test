package order

import (
	"github.com/Temich14/cart_test/internal/delivery/http/handler/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Create godoc
//
//	@Summary		Создать заказ
//	@Description	Создает новый заказ основываясь на содержимом корзины
//	@Tags			order
//	@Produce		json
//	@Param			user_id	query		uint	1	"id пользователя"
//	@Success		201		{object}	entity.Order
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/order [post]
func (h *Handler) Create(c *gin.Context) {
	id, err := utils.TryGetUserID(c)
	if err != nil {
		c.Abort()
		return
	}
	order, err := h.s.CreateNewOrder(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, order)
}
