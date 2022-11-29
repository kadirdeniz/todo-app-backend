package service

import (
	"todo/model"
	"todo/repository"
)

//go:generate mockgen -source=todo_service.go -destination=../mock/mock_service.go -package=mock
type ITodoService interface {
	CreateTodo() model.Todo
}

type TodoService struct {
	repository repository.TodoRepository
}

func NewTodoService(repository repository.TodoRepository) ITodoService {
	return &TodoService{repository: repository}
}

func (s *TodoService) CreateTodo() model.Todo {
	return model.Todo{}
}
