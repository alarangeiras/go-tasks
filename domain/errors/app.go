package errors

import "go-tasks/domain/model"

type AppError interface {
	error
	Detail() *model.ErrorDetail
}
