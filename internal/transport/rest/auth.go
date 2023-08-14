package rest

import (
	"fmt"
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
func (h *Handler) signUp(c *gin.Context) {
	var input domain.SignUpInput
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
		"id":     id,
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
func (h *Handler) signIn(c *gin.Context) {
	var input domain.SignInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	asseccToken, refreshToken, err := h.services.Auth.SignIn(input)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Header("Set-Cookie", fmt.Sprintf("refresh-token=%s; HttpOnly", refreshToken))
	c.JSON(http.StatusOK, map[string]interface{}{
		"assecc-token":  asseccToken,
		"refresh-token": refreshToken,
	})
}

func (h *Handler) refresh(c *gin.Context) {
	cookie, err := c.Cookie("refresh-token")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	accessToken, refreshToken, err := h.services.Auth.RefreshTokens(cookie)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Set-Cookie", fmt.Sprintf("refresh-token=%s; HttpOnly", refreshToken))
	c.JSON(http.StatusOK, map[string]interface{}{
		"assecc-token":  accessToken,
		"refresh-token": refreshToken,
	})
}
