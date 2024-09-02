package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) userIdentification(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	if header == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "Empty header authorization"})
		return
	}
	headerPears := strings.Split(header, " ")
	if len(headerPears) != 2 || headerPears[0] != "Bearer" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid header authorization"})
		return
	}
	if len(headerPears[1]) == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "Token is empty"})
		return
	}
	userID, err := h.service.ParseToken(headerPears[1])
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid authorization: " + err.Error()})
		return
	}
	ctx.Set("userId", userID)
}
