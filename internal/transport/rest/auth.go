package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xopxe23/books-server/internal/domain"
)

// @Summary Sign Up
// @Tags Users AUTH
// @ID sign-up
// @Accept json
// @Produce json
// @Param input body domain.SignUpInput true "Sign Up Input"
// @Success 200 {integer} int
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h Handler) signUp(c *gin.Context) {
	var input domain.SignUpInput
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

// @Summary Sign In
// @Tags Users AUTH
// @ID sign-in
// @Accept json
// @Produce json
// @Param input body domain.SignInInput true "Sign In Input"
// @Success 200 {integer} string
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h Handler) signIn(c *gin.Context) {
	var input domain.SignInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.services.Auth.GenerateToken(input)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}
