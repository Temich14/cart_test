package cart

import (
	"github.com/Temich14/cart_test/internal/service/cart"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	s *cart.Service
}

func NewHandler(s *cart.Service) *Handler {
	return &Handler{s: s}
}
func (h *Handler) Register(api *gin.RouterGroup) {
	api.POST("", h.Add)
	api.DELETE("/:product_id", h.Remove)
	api.GET("", h.Get)
	api.PATCH("", h.ChangeQuantity)
}
