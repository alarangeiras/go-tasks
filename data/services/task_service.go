package services

import (
	"go-tasks/data/ports/repositories"
	"go-tasks/domain/errors"
	"go-tasks/domain/model"
	"go-tasks/domain/usecases"
)

type TaskService struct {
	repository repositories.TaskRepository
}

var _ usecases.AddTask = (*TaskService)(nil)

func NewTaskService(repository repositories.TaskRepository) *TaskService {
	return &TaskService{
		repository,
	}
}

func (t *TaskService) AddTask(task *model.Task) errors.AppError {
	return t.repository.Add(task)
}
