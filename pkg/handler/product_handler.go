package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProducts(c *gin.Context) {
	d, err := h.services.ProductI.GetAll()
	fmt.Println(d)
	fmt.Println(err)
}
