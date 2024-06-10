package repositories

import (
	"app/domain/model"
	"app/domain/repository"
	"context"

	"gorm.io/gorm"
)

type ToDoRepository struct {
	DB *gorm.DB
}

func NewToDoRepository(db *gorm.DB) repository.ToDoRepository {
	return &ToDoRepository{DB: db}
}

func (r *ToDoRepository) GetAll(ctx context.Context) ([]*model.Todo, error) {
	var todos []*model.Todo
	if err := r.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *ToDoRepository) GetByID(ctx context.Context, id int) (*model.Todo, error) {
	var todo model.Todo
	if err := r.DB.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *ToDoRepository) Create(ctx context.Context, todo *model.Todo) error {
	return r.DB.Create(todo).Error
}

func (r *ToDoRepository) Update(ctx context.Context, todo *model.Todo) error {
	return r.DB.WithContext(ctx).Save(todo).Error
}

func (r *ToDoRepository) Delete(ctx context.Context, id int) error {
	return r.DB.Delete(&model.Todo{}, id).Error
}
