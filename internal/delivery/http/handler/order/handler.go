package order

import (
	"github.com/Temich14/cart_test/internal/service/order"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	s *order.Service
}

func NewHandler(s *order.Service) *Handler {
	return &Handler{s: s}
}
func (h *Handler) Register(api *gin.RouterGroup) {
	api.POST("", h.Create)
	api.GET("", h.GetAll)
	api.GET("/:order_id", h.GetOrder)
	api.PATCH("", h.ChangeStatus)
}
