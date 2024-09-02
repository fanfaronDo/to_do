package main

import (
	"database/sql"
	"fmt"
	"github.com/fanfaronDo/to_do/internal/config"
	"github.com/fanfaronDo/to_do/internal/domain"
	"github.com/fanfaronDo/to_do/internal/repository"
	"time"
)

func main() {
	cfg, _ := config.ConfigLoad()
	conn, err := repository.NewPostgres(cfg.Postgres)
	if err != nil {
		fmt.Println(err)
	}
	repo := repository.NewRepository(conn)

	//uid, err := repo.CreateUser(domain.User{
	//	Name:     "Herawd",
	//	Username: "heroiasdm",
	//	Password: "123",
	//})
	//if err != nil {
	//	panic(err)
	//}

	//ss, err := repo.TodoRepository.CreateItem(2, domain.TodoItem{
	//	Title:       "second",
	//	Description: sql.NullString{String: "this is second description"},
	//	DueDate:     time.Now(),
	//	CreatedAt:   time.Now(),
	//})

	ss, err := repo.TodoRepository.UpdateItem(2, 5, domain.TodoItem{
		Title:       "awd",
		Description: sql.NullString{String: "awdwaDwd"},
		DueDate:     time.Now(),
		CreatedAt:   time.Now(),
	})

	//err = repo.TodoRepository.DeleteItem(1, 2)

	fmt.Println(ss, err)
}
