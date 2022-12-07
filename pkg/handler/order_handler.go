package handler

import (
	"course_work/pkg/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllOrders(c *gin.Context) {
	orders, err := h.services.OrderI.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, orders)
}
func (h *Handler) CreateOrder(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	var cards model.UserCard
	if err := c.BindJSON(&cards); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	fmt.Println(cards)

	cards.UserId = id
	if err := h.services.OrderI.CreateOrder(cards); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
