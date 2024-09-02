package service

import (
	"github.com/fanfaronDo/to_do/internal/domain"
	"github.com/fanfaronDo/to_do/internal/repository"
)

type Todo struct {
	repo *repository.Repository
}

func NewTodo(repo *repository.Repository) *Todo {
	return &Todo{repo}
}

func (r *Todo) CreateItem(userID int, item domain.TodoItem) (domain.TodoItem, error) {
	return r.repo.TodoRepository.CreateItem(userID, item)
}

func (r *Todo) GetByItemID(userID, itemID int) (domain.TodoItem, error) {
	return r.repo.TodoRepository.GetByItemID(userID, itemID)
}

func (r *Todo) UpdateItem(userID, itemID int, item domain.TodoItem) (domain.TodoItem, error) {
	return r.repo.TodoRepository.UpdateItem(userID, itemID, item)
}

func (r *Todo) GetTodoItems(userID int) ([]domain.TodoItem, error) {
	return r.repo.TodoRepository.GetTodoItems(userID)
}

func (r *Todo) DeleteItem(userID int, itemID int) error {
	return r.repo.TodoRepository.DeleteItem(userID, itemID)
}
