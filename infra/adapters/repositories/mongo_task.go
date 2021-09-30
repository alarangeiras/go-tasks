package repositories

import (
	"go-tasks/data/ports/repositories"
	"go-tasks/domain/errors"
	"go-tasks/domain/model"
)

type MongoTaskRepository struct {
}

var _ repositories.TaskRepository = (*MongoTaskRepository)(nil)

func NewTaskRepository() repositories.TaskRepository {
	return &MongoTaskRepository{}
}

func (t MongoTaskRepository) Add(task *model.Task) errors.AppError {
	panic("implement me")
}
