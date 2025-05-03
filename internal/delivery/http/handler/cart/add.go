package cart

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	userID, err := h.tryGetUserID(c)
	if err != nil {
		c.Abort()
		return
	}
	err = h.s.AddProductToCart(userID, addDTO.ProductID, addDTO.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Product added to cart", "product_id": addDTO.ProductID})
}
