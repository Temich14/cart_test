package order

import (
	"github.com/Temich14/cart_test/internal/domain/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type statusDTO struct {
	orderID uint
	status  entity.OrderStatus
}

func (h *Handler) ChangeStatus(c *gin.Context) {
	var dto statusDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.s.ChangeStatus(dto.orderID, dto.status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": dto.status})
}
