package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetProducts(c *gin.Context) {
	products, err := h.services.ProductI.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, products)
}
func (h *Handler) GetAllCategories(c *gin.Context) {
	categories, err := h.services.ProductI.GetAllCategories()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, categories)
}
func (h *Handler) GetAllProviders(c *gin.Context) {
	providers, err := h.services.ProductI.GetAllProviders()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, providers)
}
