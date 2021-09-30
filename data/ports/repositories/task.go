package repositories

import (
	"go-tasks/domain/errors"
	"go-tasks/domain/model"
)

type TaskRepository interface {
	Add(task *model.Task) errors.AppError
}
