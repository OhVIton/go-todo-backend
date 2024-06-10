package repository

import (
	"app/domain/model"
	"context"
)

type ToDoRepository interface {
	GetByID(ctx context.Context, id int) (*model.Todo, error)
	GetAll(ctx context.Context) ([]*model.Todo, error)
	Create(ctx context.Context, todo *model.Todo) error
	Update(ctx context.Context, todo *model.Todo) error
	Delete(ctx context.Context, id int) error
}
