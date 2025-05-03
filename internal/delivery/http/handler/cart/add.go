package cart

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AddDTO struct {
	ProductID uint
	Quantity  int
}

func (h *Handler) Add(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.AbortWithStatus(http.StatusMethodNotAllowed)
		return
	}
	var addDTO AddDTO
	if err := c.ShouldBindJSON(&addDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userIDStr, exists := c.Get("user_id")
	if !exists {
		if userIDStr = c.Query("user_id"); userIDStr != "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No user_id found"})
			return
		}
	}
	userID, err := strconv.ParseUint(userIDStr.(string), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.s.AddProductToCart(uint(userID), addDTO.ProductID, addDTO.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Product added to cart", "product_id": addDTO.ProductID})
}
