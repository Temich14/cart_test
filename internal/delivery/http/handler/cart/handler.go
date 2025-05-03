package cart

import (
	"errors"
	"github.com/Temich14/cart_test/internal/service/cart"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
func (h *Handler) tryGetUserID(c *gin.Context) (uint, error) {
	userIDStr, exists := c.Get("user_id")
	if !exists {
		if userIDStr = c.Query("user_id"); userIDStr != "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found"})
			return -1, errors.New("user_id not found")
		}
	}
	userID, err := strconv.ParseUint(userIDStr.(string), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return -1, err
	}
	return uint(userID), nil
}
