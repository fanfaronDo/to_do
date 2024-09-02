package handler

import (
	"errors"
	"github.com/fanfaronDo/to_do/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func (h *Handler) getTodoItemsById(ctx *gin.Context) {
	userID, err := getUserId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user error: " + err.Error()})
		return
	}
	itemId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found id: " + err.Error()})
		return
	}
	itemResp, err := h.service.TodoService.GetByItemID(userID, itemId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, itemResp)
}

func (h *Handler) updateTodoItem(ctx *gin.Context) {
	userID, err := getUserId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user error: " + err.Error()})
		return
	}
	itemID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found id: " + err.Error()})
		return
	}
	var item domain.TodoItem
	if err = ctx.BindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad input data format: " + err.Error()})
		return
	}

	itemResp, err := h.service.UpdateItem(userID, itemID, item)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &itemResp)
}

func (h *Handler) deleteItem(ctx *gin.Context) {
	userID, err := getUserId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user error: " + err.Error()})
		return
	}
	itemID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found id: " + err.Error()})
		return
	}
	err = h.service.DeleteItem(userID, itemID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
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
