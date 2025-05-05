package cart

import (
	"github.com/Temich14/cart_test/internal/delivery/http/handler/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Get godoc
//
//	@Summary		Получить корзину пользователя
//	@Description	Возвращает корзину пользователя со списком товаров
//	@Tags			cart
//	@Produce		json
//	@Param			page	query		int			false	"страница пагинации"	example(1)
//	@Param			limit	query		int			false	"лимит пагинации"		example(10)
//	@Param			user_id	query		uint	false	"id пользователя"	example(1)
//	@Success		200	{object}	entity.Cart
//	@Failure		500	{object}	map[string]string
//	@Router			/cart [get]
func (h *Handler) Get(c *gin.Context) {
	userID, err := utils.TryGetUserID(c)
	if err != nil {
		c.Abort()
		return
	}
	pageStr := c.Query("page")
	limitStr := c.Query("limit")
	page := 0
	limit := 0
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
	cart, err := h.s.GetUserCart(userID, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cart)
}
