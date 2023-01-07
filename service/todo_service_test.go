package service

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"todo/mock"
	"todo/model"
)

func TestService_CreateTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	todoObj := model.Todo{
		ID:   "123",
		Text: "buy some milk",
	}

	mockRepository := mock.NewMockITodoRepository(ctrl)
	mockRepository.
		EXPECT().
		CreateTodo(
			gomock.Any(),
			gomock.Any(),
		).
		Return(todoObj).
		Times(1)

	service := NewTodoService(mockRepository)
	todo := service.CreateTodo("buy some milk")

	assert.Equal(t, todoObj.Text, todo.Text)
}
