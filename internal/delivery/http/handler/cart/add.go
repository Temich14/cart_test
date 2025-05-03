package cart

import (
	"github.com/Temich14/cart_test/internal/delivery/http/handler/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type addDTO struct {
	ProductID uint
	Quantity  int
}

func (h *Handler) Add(c *gin.Context) {
	var dto addDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, err := utils.TryGetUserID(c)
	if err != nil {
		c.Abort()
		return
	}
	err = h.s.AddProductToCart(userID, dto.ProductID, dto.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Product added to cart", "product_id": dto.ProductID})
}
