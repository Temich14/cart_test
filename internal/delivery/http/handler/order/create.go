package order

import (
	"github.com/Temich14/cart_test/internal/delivery/http/handler/utils"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"runtime/debug"
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
		h.log.Error(
			"error getting user id",
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return
	}
	h.log.Debug("creating new order", slog.Uint64("user_id", uint64(id)))
	order, err := h.s.CreateNewOrder(id)
	if err != nil {
		h.log.Error("error creating new order", slog.Uint64("user_id", uint64(id)), slog.String("error", err.Error()), slog.String("stack", string(debug.Stack())))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	h.log.Debug("order created", slog.Uint64("user_id", uint64(id)), slog.Uint64("order_id", uint64(order.ID)))
	c.JSON(http.StatusCreated, order)
}
