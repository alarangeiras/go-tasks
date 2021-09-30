package errors

import "go-tasks/domain/model"

type baseError struct {
	detail *model.ErrorDetail
}

var _ AppError = (*baseError)(nil)

func (b *baseError) Error() string {
	return b.detail.Message
}

func (b *baseError) Detail() *model.ErrorDetail {
	return b.detail
}

func newBaseError(code int, message string) AppError {
	return &baseError{
		&model.ErrorDetail{
			code,
			message,
		},
	}
}
