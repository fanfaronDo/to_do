package handler

import (
	"github.com/fanfaronDo/to_do/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var user domain.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid input: " + err.Error()})
		return
	}
	id, err := h.service.AuthorizationService.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "User should be not created: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"User": id})
}

func (h *Handler) signIn(ctx *gin.Context) {
	tmpUser := struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}

	if err := ctx.BindJSON(&tmpUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid input: " + err.Error()})
		return
	}
	token, err := h.service.AuthorizationService.GenerateToken(tmpUser.Username, tmpUser.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Token should be not generated: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
