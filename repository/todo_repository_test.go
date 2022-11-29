package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"todo/model"
)

func TestRepository_CreateTodo(t *testing.T) {
	todoList := []model.Todo{}

	repository := NewTodoRepository()
	todo := repository.CreateTodo(
		&todoList,
		model.Todo{
			ID:   "123",
			Text: "test",
		},
	)

	assert.Equal(t, todoList[0], todo)
}
