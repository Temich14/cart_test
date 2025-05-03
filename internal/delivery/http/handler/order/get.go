package order

import (
	"github.com/Temich14/cart_test/internal/delivery/http/handler/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetAll(c *gin.Context) {
	userID, err := utils.TryGetUserID(c)
	if err != nil {
		c.Abort()
		return
	}
	orders, err := h.s.GetOrders(userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, orders)
}
func (h *Handler) GetOrder(c *gin.Context) {
	orderIDStr := c.Param("orderID")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 32)
	if err != nil {
		c.Abort()
		return
	}
	order, err := h.s.GetOrder(uint(orderID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, order)
}
