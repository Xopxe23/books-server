package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xopxe23/books-server/internal/domain"
)

func (h Handler) signUp(c *gin.Context) {
	var input domain.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Auth.SignUp(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
		"status": "User created",
	})
}
func (h Handler) signIn(c *gin.Context) {}
