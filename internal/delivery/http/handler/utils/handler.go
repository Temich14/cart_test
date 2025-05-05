package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func TryGetUserID(c *gin.Context) (uint, error) {
	userIDStr, exists := c.Get("user_id")
	if !exists {
		if userIDStr = c.Query("user_id"); userIDStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found"})
			return 0, errors.New("user_id not found")
		}
	}
	userID, err := strconv.ParseUint(userIDStr.(string), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0, err
	}
	return uint(userID), nil
}
