package cart

import (
	"github.com/Temich14/cart_test/internal/delivery/http/handler/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Remove godoc
//
//	@Summary		Удалить товар из корзины
//	@Description	Удаляет товар с указанным ID из корзины пользователя
//	@Tags			cart
//	@Produce		json
//	@Param			user_id	query		uint		false	"id пользователя"	example(1)
//	@Param			product_id	path		int	true	"ID товара"
//	@Success		200			{object}	entity.Cart
//	@Failure		400			{object}	map[string]string
//	@Failure		500			{object}	map[string]string
//	@Router			/cart/{product_id} [delete]
func (h *Handler) Remove(c *gin.Context) {
	idStr := c.Param("product_id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, err := utils.TryGetUserID(c)
	if err != nil {
		c.Abort()
		return
	}
	productID, err := h.s.RemoveProductFromCart(userID, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cart": productID})
}
