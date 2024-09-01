package repository

import (
	"database/sql"
	"fmt"
	"github.com/fanfaronDo/to_do/internal/config"
	_ "github.com/lib/pq"
)

func NewPostgres(cfg config.Postgres) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Database, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
