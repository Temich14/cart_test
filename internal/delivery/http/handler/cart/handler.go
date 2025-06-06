package cart

import (
	"github.com/gin-gonic/gin"
	"log/slog"
)

type Handler struct {
	s   CartService
	log *slog.Logger
}

func NewHandler(s CartService, log *slog.Logger) *Handler {
	return &Handler{s: s, log: log}
}
func (h *Handler) Register(api *gin.RouterGroup) {
	api.POST("", h.Add)
	api.DELETE("/:product_id", h.Remove)
	api.GET("", h.Get)
	api.PATCH("", h.ChangeQuantity)
}
