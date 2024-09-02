package service

import (
	"fmt"
	"github.com/fanfaronDo/to_do/internal/domain"
	"github.com/fanfaronDo/to_do/internal/repository"
	"time"
)

type Todo struct {
	repo *repository.Repository
}

func NewTodo(repo *repository.Repository) *Todo {
	return &Todo{repo}
}

func (s *Todo) CreateItem(userID int, item domain.TodoItem) (domain.TodoItem, error) {
	item.CreatedAt = time.Now().Format(time.RFC3339)
	item.UpdatedAt = time.Now().Format(time.RFC3339)
	fmt.Println(item)
	return s.repo.TodoRepository.CreateItem(userID, item)
}

func (s *Todo) GetByItemID(userID, itemID int) (domain.TodoItem, error) {
	return s.repo.TodoRepository.GetByItemID(userID, itemID)
}

func (s *Todo) UpdateItem(userID, itemID int, item domain.TodoItem) (domain.TodoItem, error) {
	item.UpdatedAt = time.Now().Format(time.RFC3339)
	return s.repo.TodoRepository.UpdateItem(userID, itemID, item)
}

func (s *Todo) GetTodoItems(userID int) ([]domain.TodoItem, error) {
	return s.repo.TodoRepository.GetTodoItems(userID)
}

func (s *Todo) DeleteItem(userID int, itemID int) error {
	return s.repo.TodoRepository.DeleteItem(userID, itemID)
}

//func (s *Todo) parseTime(dateString string) (time.Time, error) {
//	layout := "2006-01-02 15:04:05"
//	t, err := time.Parse(layout, dateString)
//	if err != nil {
//		return time.Time{}, err
//	}
//	return t, nil
//}
