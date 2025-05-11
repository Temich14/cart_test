package cart

import (
	"fmt"
	"github.com/Temich14/cart_test/internal/delivery/http/handler/utils"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"runtime/debug"
	"strconv"
)

// Remove godoc
//
//	@Summary		Удалить товар из корзины
//	@Description	Удаляет товар с указанным ID из корзины пользователя
//	@Tags			cart
//	@Produce		json
//	@Param			user_id		query		uint	false	"id пользователя"	example(1)
//	@Param			product_id	path		int		true	"ID товара"
//	@Success		200			{object}	entity.Cart
//	@Failure		400			{object}	map[string]string
//	@Failure		500			{object}	map[string]string
//	@Router			/cart/{product_id} [delete]
func (h *Handler) Remove(c *gin.Context) {
	idStr := c.Param("product_id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.log.Error(
			"error parsing product id",
			slog.String("product_id string", idStr),
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, err := utils.TryGetUserID(c)
	if err != nil {
		h.log.Error(
			"error while trying to get user ID",
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return
	}
	h.log.Debug(
		"removing product from cart",
		slog.Int("product_id", int(id)),
		slog.Int("user_id", int(userID)))
	productID, err := h.s.RemoveProductFromCart(userID, uint(id))
	if err != nil {
		h.log.Error(
			"error removing product from cart",
			slog.Int("user_id", int(userID)),
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmtDebug := fmt.Sprintf("product %d was removed from cart", int(id))
	h.log.Debug(fmtDebug, slog.Int("user_id", int(userID)))
	c.JSON(http.StatusOK, gin.H{"cart": productID})
}
