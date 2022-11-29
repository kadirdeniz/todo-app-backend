package service

import (
	"math/rand"
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
	return model.Todo{
		ID:   string(rand.Intn(100)),
		Text: todo,
	}
}
