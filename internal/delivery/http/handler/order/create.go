package order

import (
	"github.com/Temich14/cart_test/internal/domain/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type cartDTO struct {
	cart entity.Cart
}

func (h *Handler) Create(c *gin.Context) {
	dto := &cartDTO{}
	if err := c.ShouldBindJSON(dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	order, err := h.s.CreateNewOrder(&dto.cart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, order)
}
