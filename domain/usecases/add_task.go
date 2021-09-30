package usecases

import (
	"go-tasks/domain/errors"
	"go-tasks/domain/model"
)

type AddTask interface {
	AddTask(task *model.Task) errors.AppError
}
