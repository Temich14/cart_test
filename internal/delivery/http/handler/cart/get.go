package cart

import (
	"github.com/Temich14/cart_test/internal/delivery/http/handler/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Get(c *gin.Context) {
	userID, err := utils.TryGetUserID(c)
	if err != nil {
		c.Abort()
		return
	}
	cart, err := h.s.GetUserCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cart)
}
