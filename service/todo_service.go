package service

import (
	"github.com/google/uuid"
	"todo/model"
	"todo/repository"
)

//go:generate mockgen -source=todo_service.go -destination=../mock/mock_service.go -package=mock
type ITodoService interface {
	CreateTodo(todo string) model.Todo
}

type TodoService struct {
	repository repository.ITodoRepository
}

func NewTodoService(repository repository.ITodoRepository) ITodoService {
	return &TodoService{repository: repository}
}

func (s *TodoService) CreateTodo(todo string) model.Todo {
	todoObj := model.Todo{
		ID:   uuid.New().String(),
		Text: todo,
	}
	return s.repository.CreateTodo(
		&repository.TodoList,
		todoObj,
	)
}
