package services

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go-tasks/data/ports/repositories/mock"
	"go-tasks/domain/errors"
	"go-tasks/domain/model"
	"go-tasks/domain/usecases"
	"testing"
)

func TestTaskService_AddTask_Ok(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	taskRepository := mock.NewMockTaskRepository(ctrl)
	taskRepository.EXPECT().Add(gomock.Any()).Return(nil)

	var usecase usecases.AddTask = NewTaskService(taskRepository)
	task := &model.Task{}
	appError := usecase.AddTask(task)

	assert.Nil(t, appError)
}

func TestTaskService_AddTask_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	taskRepository := mock.NewMockTaskRepository(ctrl)
	taskRepository.EXPECT().Add(gomock.Any()).Return(errors.NewNotFound())

	var usecase usecases.AddTask = NewTaskService(taskRepository)
	task := &model.Task{}
	appError := usecase.AddTask(task)

	assert.NotNil(t, appError)
}
