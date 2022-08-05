package core

import (
	"context"
	"github.com/google/uuid"
	"github.com/todolist-project-rhttraining/src/todoservice/pkg/repository"
	"github.com/todolist-project-rhttraining/src/todoservice/pkg/repository/db"
)

type TodoService struct {
	repo repository.IRepository
}

func NewTodoService(repo repository.IRepository) TodoService {
	return TodoService{
		repo: repo,
	}
}

// CreateTodo return string object created
func (ts TodoService) CreateTodo(ctx context.Context, strAuthId string, todo string, done bool) (string, error) {
	userId, err := uuid.Parse(strAuthId)
	if err != nil {
		return "", err
	}

	id, err := ts.repo.Todo().InsertList(ctx, userId, todo, done)
	return id.String(), err
}

func (ts TodoService) GetTodo(ctx context.Context, strAuthId string) ([]db.List, error) {
	userId, err := uuid.Parse(strAuthId)
	if err != nil {
		return []db.List{}, err
	}

	return ts.repo.Todo().GetList(ctx, userId)
}
