package services

import (
	"app/domain/model"
	"app/domain/repository"
	"context"
)

type ToDoService struct {
	repo repository.ToDoRepository
}

func NewToDoService(repo repository.ToDoRepository) *ToDoService {
	return &ToDoService{repo: repo}
}

func (s *ToDoService) GetAll(ctx context.Context) ([]*model.Todo, error) {
	return s.repo.GetAll(ctx)
}

func (s *ToDoService) GetByID(ctx context.Context, id int) (*model.Todo, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ToDoService) Create(ctx context.Context, content string) error {
	todo := &model.Todo{Content: content}
	return s.repo.Create(ctx, todo)
}

func (s *ToDoService) Update(ctx context.Context, todo *model.Todo) error {
	return s.repo.Update(ctx, todo)
}

func (s *ToDoService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
