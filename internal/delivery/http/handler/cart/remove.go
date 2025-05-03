package cart

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) Remove(c *gin.Context) {
	if c.Request.Method == http.MethodDelete {
		c.Status(http.StatusMethodNotAllowed)
		return
	}
	idStr := c.Param("product_id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, err := h.tryGetUserID(c)
	if err != nil {
		c.Abort()
		return
	}
	err = h.s.RemoveProductFromCart(userID, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product removed", "product_id": id})
}
