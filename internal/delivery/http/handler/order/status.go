package order

import (
	"github.com/Temich14/cart_test/internal/domain/entity"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"runtime/debug"
)

type statusDTO struct {
	OrderID uint               `json:"order_id" example:"1"`
	Status  entity.OrderStatus `json:"status" example:"completed"`
}

// ChangeStatus godoc
//
//	@Summary		Изменить статус заказа
//	@Description	Изменяет статус указанного заказа
//	@Tags			order
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query		uint				false	"id пользователя"	example(1)
//	@Param			input	body		statusDTO			true	"ID заказа и новый статус"
//	@Success		200		{object}	entity.Order
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/order/status [patch]
func (h *Handler) ChangeStatus(c *gin.Context) {
	var dto statusDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		h.log.Error("error binding json", slog.String("err", err.Error()), slog.String("stack", string(debug.Stack())))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.log.Debug("changing status", slog.Uint64("order_id", uint64(dto.OrderID)))
	order, err := h.s.ChangeStatus(dto.OrderID, dto.Status)
	if err != nil {
		h.log.Error(
			"error changing status",
			slog.Uint64("order_id", uint64(dto.OrderID)),
			slog.String("err", err.Error()),
			slog.String("stack", string(debug.Stack())))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	h.log.Debug("status changed", slog.Uint64("order_id", uint64(dto.OrderID)), slog.String("new_status", order.Status))
	c.JSON(http.StatusOK, gin.H{"order": order})
}
