package repository

import (
	"database/sql"
	"github.com/fanfaronDo/to_do/internal/domain"
)

type Authorization struct {
	db *sql.DB
}

func NewAuthorization(db *sql.DB) *Authorization {
	return &Authorization{db: db}
}

func (a *Authorization) CreateUser(user domain.User) (int, error) {
	var id int
	query := "INSERT INTO users (name, username, password) VALUES ($1, $2, $3) RETURNING user_id"
	row := a.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (a *Authorization) GetUser(username, password string) (domain.User, error) {
	var user domain.User
	query := "SELECT user_id, name, username, password FROM users WHERE username = $1 AND password = $2"
	row := a.db.QueryRow(query, username, password)
	err := row.Scan(&user.Id, &user.Name, &user.Username, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}
