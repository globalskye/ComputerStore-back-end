package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetCustomers(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": userId})
}

func (h *Handler) GetCustomer(c *gin.Context) {

}
func (h *Handler) CreateCustomer(c *gin.Context) {

}
func (h *Handler) UpdateCustomer(c *gin.Context) {

}
func (h *Handler) DeleteCustomer(c *gin.Context) {

}
