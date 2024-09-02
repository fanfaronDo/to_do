package repository

import (
	"database/sql"
	"github.com/fanfaronDo/to_do/internal/domain"
)

type AuthorisationRepository interface {
	CreateUser(user domain.User) (int, error)
	GetUser(username, password string) (domain.User, error)
}

type TodoRepository interface {
	CreateItem(userID int, item domain.TodoItem) (domain.TodoItem, error)
	GetTodoItems(userID int) ([]domain.TodoItem, error)
	GetByItemID(userID, itemID int) (domain.TodoItem, error)
	UpdateItem(userID, itemID int, item domain.TodoItem) (domain.TodoItem, error)
	DeleteItem(userID, itemID int) error
}

type Repository struct {
	AuthorisationRepository
	TodoRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		AuthorisationRepository: NewAuthorization(db),
		TodoRepository:          NewTodoItemPg(db),
	}
}
