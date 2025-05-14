package order

import (
	"github.com/gin-gonic/gin"
	"log/slog"
)

type Handler struct {
	s   OrderService
	log *slog.Logger
}

func NewHandler(s OrderService, log *slog.Logger) *Handler {
	return &Handler{s: s, log: log}
}
func (h *Handler) Register(api *gin.RouterGroup) {
	api.POST("", h.Create)
	api.GET("", h.GetAll)
	api.GET("/:order_id", h.GetOrder)
	api.PATCH("", h.ChangeStatus)
}
