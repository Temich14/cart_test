package order

import (
	"fmt"
	"github.com/Temich14/cart_test/internal/delivery/http/handler/utils"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"runtime/debug"
	"strconv"
)

// GetAll godoc
//
//	@Summary		Получить заказы пользователя
//	@Description	Возвращает список всех заказов пользователя
//	@Tags			order
//	@Produce		json
//	@Param			user_id	query		uint		false	"id пользователя"	example(1)
//	@Param			status	query		string		false	"сортировка по статуса"	example(completed)
//	@Param			page	query		int			false	"страница пагинации"	example(1)
//	@Param			limit	query		int			false	"лимит пагинации"		example(10)
//	@Success		200		{array}		entity.Order
//	@Failure		500		{object}	map[string]string
//	@Router			/order 	[get]
func (h *Handler) GetAll(c *gin.Context) {
	userID, err := utils.TryGetUserID(c)
	if err != nil {
		h.log.Error(
			"error getting user id",
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return
	}
	status := c.Query("status")
	pageStr := c.Query("page")
	limitStr := c.Query("limit")
	page := 0
	limit := 0
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			h.log.Error(
				"error parsing page",
				slog.Uint64("user_id", uint64(userID)),
				slog.String("error", err.Error()),
				slog.String("stack", string(debug.Stack())))
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			h.log.Error(
				"error parsing limit",
				slog.Uint64("user_id", uint64(userID)),
				slog.String("error", err.Error()),
				slog.String("stack", string(debug.Stack())))
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	fmtMsg := fmt.Sprintf("getting list of orders in page %d and limit = %d", page, limit)
	h.log.Debug(fmtMsg, slog.Uint64("user_id", uint64(userID)))
	orders, err := h.s.GetOrders(userID, status, page, limit)
	if err != nil {
		h.log.Error(
			"error getting orders",
			slog.Uint64("user_id", uint64(userID)),
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return
	}
	h.log.Debug("got list of orders", slog.Uint64("user_id", uint64(userID)))
	c.JSON(http.StatusOK, orders)
}

// GetOrder godoc
//
//	@Summary		Получить заказ по ID
//	@Description	Возвращает заказ с указанным ID
//	@Tags			order
//	@Produce		json
//	@Param			user_id	query		uint		false	"id пользователя"	example(1)
//	@Param			orderID	path		int	true	"ID заказа"
//	@Success		200		{object}	entity.Order
//	@Failure		400		{object}	map[string]string
//	@Router			/order/{orderID} [get]
func (h *Handler) GetOrder(c *gin.Context) {
	orderIDStr := c.Param("order_id")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 32)
	if err != nil {
		h.log.Error(
			"error getting user id",
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		return
	}
	fmtMsg := fmt.Sprintf("getting order by id %d", orderID)
	h.log.Debug(fmtMsg, slog.Uint64("order_id", orderID))
	order, err := h.s.GetOrder(uint(orderID))
	if err != nil {
		h.log.Error(
			"error getting order",
			slog.Uint64("order_id", orderID),
			slog.String("error", err.Error()),
			slog.String("stack", string(debug.Stack())))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.log.Debug("got order", slog.Uint64("order_id", orderID))
	c.JSON(http.StatusOK, order)
}
