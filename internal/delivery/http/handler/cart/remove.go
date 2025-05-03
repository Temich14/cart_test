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
	userIDStr, exists := c.Get("user_id")
	if !exists {
		if userIDStr = c.Query("user_id"); userIDStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No user_id found"})
			return
		}
	}
	userID, err := strconv.ParseUint(userIDStr.(string), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.s.RemoveProductFromCart(uint(userID), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product removed", "product_id": id})
}
