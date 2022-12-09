package handler

import (
	"course_work/pkg/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllOrders(c *gin.Context) {
	orders, err := h.services.OrderI.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, orders)
	return
}
func (h *Handler) CreateOrder(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var cards model.UserCard
	if err := c.Bind(&cards); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		fmt.Println(err)
		return
	}
	if len(cards.Items) < 1 {
		newErrorResponse(c, http.StatusBadRequest, errors.New("Items cannot be null").Error())
		return
	}

	cards.UserId = id
	if err := h.services.OrderI.CreateOrder(cards); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
	return
}
