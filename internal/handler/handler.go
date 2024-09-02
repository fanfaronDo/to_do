package handler

import (
	"github.com/fanfaronDo/to_do/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign_up", h.signUp)
		auth.POST("/sign_in", h.signIn)
	}
	api := router.Group("/api", h.userIdentification)
	{
		tasks := api.Group("/tasks")
		{
			tasks.POST("/", h.createTodoItem)
		}
	}

	return router
}
