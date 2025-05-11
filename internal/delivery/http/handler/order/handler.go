package order

import (
	"github.com/Temich14/cart_test/internal/domain/service/order"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type Handler struct {
	s   *order.Service
	log *slog.Logger
}

func NewHandler(s *order.Service, log *slog.Logger) *Handler {
	return &Handler{s: s, log: log}
}
func (h *Handler) Register(api *gin.RouterGroup) {
	api.POST("", h.Create)
	api.GET("", h.GetAll)
	api.GET("/:order_id", h.GetOrder)
	api.PATCH("", h.ChangeStatus)
}
