package handler

import (
	"errors"
	"github.com/fanfaronDo/to_do/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createTodoItem(ctx *gin.Context) {
	var item domain.TodoItem
	userID, err := getUserId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user error: " + err.Error()})
		return
	}

	if err = ctx.BindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad input data format: " + err.Error()})
		return
	}

	todoItemResp, err := h.service.CreateItem(userID, item)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, &todoItemResp)
}

func (h *Handler) getTodoItems(ctx *gin.Context) {
	userID, err := getUserId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user error: " + err.Error()})
		return
	}
	todoItems, err := h.service.GetTodoItems(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, todoItems)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get("userId")
	if !ok {
		return 0, errors.New("User id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("User id is of invalid type")
	}

	return idInt, nil
}
