package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "auth header left")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	if headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "invalid token type, should be bearer")
		return
	}
	userId, err := h.services.Authorization.ParseAccessToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	user, err := h.services.Authorization.GetUserById(userId)
	if len(user) > 0 {
		c.Set("role", user[0].Role)
		c.Set("userId", userId)
		c.Next()
	} else {
		newErrorResponse(c, http.StatusUnauthorized, errors.New("Token exists, but user in database not found").Error())
		return
	}

	return
}

func (h *Handler) adminIdentity(c *gin.Context) {
	if role, _ := c.Get("role"); role != "admin" {
		newErrorResponse(c, http.StatusUnauthorized, "U dont have access, u not a admin")
		return
	}
	c.Next()
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get("userId")
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
