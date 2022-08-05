package repository

import (
	"github.com/todolist-project-rhttraining/src/todoservice/config"
	"github.com/todolist-project-rhttraining/src/todoservice/pkg/repository/db"
)

type IRepository interface {
}

type Repository struct {
	dbSource db.DBSource
}

func NewRepository(cfg config.DbConfig) (Repository, error) {
	dbSource, err := db.NewDatabaseSource(cfg)
	if err != nil {
		return Repository{}, err
	}

	return Repository{
		dbSource: dbSource,
	}, nil
}
