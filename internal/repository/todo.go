package repository

import (
	"database/sql"
	"github.com/fanfaronDo/to_do/internal/domain"
	"time"
)

type TodoItemPg struct {
	db *sql.DB
}

func NewTodoItemPg(db *sql.DB) *TodoItemPg {
	return &TodoItemPg{db: db}
}

func (t *TodoItemPg) CreateItem(userID int, item domain.TodoItem) (domain.TodoItem, error) {
	tr, err := t.db.Begin()
	if err != nil {
		return domain.TodoItem{}, err
	}
	queryInsertItem := "INSERT INTO todo_items(title, description, due_date, created_at) VALUES ($1,$2,$3,$4) RETURNING id"
	row := tr.QueryRow(queryInsertItem, item.Title, item.Description.String, item.DueDate, item.CreatedAt)
	err = row.Scan(&item.ID)
	if err != nil {
		tr.Rollback()
		return domain.TodoItem{}, err
	}
	queryInsertUserTodoItems := "INSERT INTO user_todo_items(user_id, todo_id) VALUES ($1, $2)"
	_, err = tr.Exec(queryInsertUserTodoItems, userID, item.ID)

	if err != nil {
		tr.Rollback()
		return domain.TodoItem{}, err
	}

	err = tr.Commit()

	if err != nil {
		tr.Rollback()
		return domain.TodoItem{}, err
	}

	return item, nil
}

func (t *TodoItemPg) GetByItemID(userID, itemID int) (domain.TodoItem, error) {
	var todoItem domain.TodoItem
	query := "SELECT ti.id, title, description, due_date, created_at, updated_at " +
		"FROM todo_items ti JOIN user_todo_items uti ON ti.id = uti.todo_id " +
		"WHERE ti.id = $1 AND uti.user_id = $2"
	row := t.db.QueryRow(query, itemID, userID)
	err := row.Scan(&todoItem.ID, &todoItem.Title, &todoItem.Description, &todoItem.DueDate, &todoItem.CreatedAt, &todoItem.UpdatedAt)
	if err != nil {
		return domain.TodoItem{}, err
	}
	return todoItem, nil
}

func (t *TodoItemPg) DeleteItem(userID, itemID int) error {
	tr, err := t.db.Begin()
	if err != nil {
		return err
	}
	queryDelUserItem := "DELETE FROM user_todo_items WHERE todo_id = $1 and user_id = $2"
	_, err = tr.Exec(queryDelUserItem, itemID, userID)
	if err != nil {
		tr.Rollback()
		return err
	}
	queryDelItem := "DELETE FROM todo_items WHERE id = $1"
	_, err = tr.Exec(queryDelItem, itemID)
	if err != nil {
		tr.Rollback()
		return err
	}

	err = tr.Commit()

	if err != nil {
		return err
	}

	return nil
}

func (t *TodoItemPg) UpdateItem(userID, itemID int, item domain.TodoItem) (domain.TodoItem, error) {
	item.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}
	query := "UPDATE todo_items ti SET title = $1, description = $2, due_date = $3, created_at = $4, updated_at = $5 " +
		"FROM user_todo_items uti WHERE ti.id = uti.todo_id AND uti.user_id = $6 AND ti.id = $7;"
	_, err := t.db.Exec(query,
		item.Title,
		item.Description.String,
		item.DueDate,
		item.CreatedAt,
		item.UpdatedAt.Time,
		userID,
		itemID)

	if err != nil {
		return domain.TodoItem{}, err
	}

	return item, nil
}

func (t *TodoItemPg) GetTodoItems(userID int) ([]domain.TodoItem, error) {
	var items []domain.TodoItem
	query := "select ti.id, title, description, due_date, created_at, updated_at " +
		"from todo_items ti left join user_todo_items ui on ti.id = ui.todo_id where user_id = $1;"
	row, err := t.db.Query(query, userID)
	if err != nil {
		return items, err
	}
	for row.Next() {
		var item domain.TodoItem
		err = row.Scan(&item.ID, &item.Title, &item.Description, &item.DueDate, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return items, err
		}
		items = append(items, item)
	}

	return items, nil
}
