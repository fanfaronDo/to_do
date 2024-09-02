package service

import (
	"github.com/fanfaronDo/to_do/internal/domain"
	"github.com/fanfaronDo/to_do/internal/repository"
)

type AuthorizationService interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type TodoService interface {
	CreateItem(userID int, item domain.TodoItem) (domain.TodoItem, error)
	GetByItemID(userID, itemID int) (domain.TodoItem, error)
	UpdateItem(userID, itemID int, item domain.TodoItem) (domain.TodoItem, error)
	GetTodoItems(userID int) ([]domain.TodoItem, error)
	DeleteItem(userID, itemID int) error
}

type Service struct {
	AuthorizationService
	TodoService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		AuthorizationService: NewAuthorization(repo),
		TodoService:          NewTodo(repo),
	}
}
