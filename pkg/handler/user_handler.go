package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.services.UserI.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handler) DeleteUserById(c *gin.Context) {

	data := c.Param("id")
	fmt.Println(data)
	id, err := strconv.Atoi(data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	if err := h.services.UserI.DeleteById(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"message": "sucsfully"})

}
