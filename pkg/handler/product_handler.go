package handler

import (
	"course_work/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetAllProducts(c *gin.Context) {
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

func (h *Handler) DeleteProduct(c *gin.Context) {
	itemId := c.Param("id")
	id, err := strconv.Atoi(itemId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.ProductI.DeleteById(id); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"messsage": "ok"})
}
func (h *Handler) PostToProducts(c *gin.Context) {
	var product model.ProductToAdd
	if err := c.BindJSON(product); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.ProductI.PostProductToStock(product); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})

}

func (h *Handler) GetAllProviders(c *gin.Context) {
	providers, err := h.services.ProductI.GetAllProviders()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, providers)
}
