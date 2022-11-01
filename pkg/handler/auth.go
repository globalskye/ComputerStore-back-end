package handler

import (
	"course_work/pkg/model"
	"course_work/pkg/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateUser(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})

}
func (h *Handler) signIn(c *gin.Context) {
	db, _ := repository.NewPostgresDb()
	row := db.QueryRow("INSERT INTO users (username, email, password_hash) VALUES($1,$2,$3) RETURNING id", "4", "4", "4")
	var id int
	row.Scan(&id)

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}
