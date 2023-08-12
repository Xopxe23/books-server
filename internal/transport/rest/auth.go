package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xopxe23/books-server/internal/domain"
)

// @Summary Create User
// @Tags Users AUTH
// @ID create-user
// @Accept json
// @Produce json
// @Param input body domain.User true "Update input"
// @Success 200 {integer} int
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h Handler) signUp(c *gin.Context) {
	var input domain.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Auth.CreateUser(input)
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
