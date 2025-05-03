package cart

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Get(c *gin.Context) {
	userID, err := h.tryGetUserID(c)
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
