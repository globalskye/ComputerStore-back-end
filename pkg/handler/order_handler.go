package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllOrders(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"wqe": "qwe"})

}
