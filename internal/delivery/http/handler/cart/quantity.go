package cart

import (
	"github.com/Temich14/cart_test/internal/delivery/http/handler/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type quantityDTO struct {
	ProductID uint `json:"product_id" example:"1"`
	Quantity  int  `json:"quantity" example:"4"`
}

// ChangeQuantity godoc
//
//	@Summary		Изменить количество товара в корзине
//	@Description	Изменяет количество указанного товара в корзине пользователя. Не добавляет товар, если его нет в корзине
//	@Tags			cart
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query		uint		false	"id пользователя"	example(1)
//	@Param			input	body		quantityDTO	true	"ID товара и новое количество"
//	@Success		200		{string}	entity.Cart
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/cart [patch]
func (h *Handler) ChangeQuantity(c *gin.Context) {
	userID, err := utils.TryGetUserID(c)
	if err != nil {
		c.Abort()
		return
	}
	var dto quantityDTO
	err = c.BindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	productID, err := h.s.ChangeQuantity(userID, dto.ProductID, dto.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"product_id": productID})
}
