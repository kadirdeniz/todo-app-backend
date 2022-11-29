package repository

import "todo/model"

//go:generate mockgen -source=todo_repository.go -destination=../mock/mock_repository.go -package=mock
type ITodoRepository interface {
	CreateTodo() model.Todo
}

type TodoRepository struct {
}

func NewTodoRepository() ITodoRepository {
	return &TodoRepository{}
}

func (r *TodoRepository) CreateTodo() model.Todo {
	return model.Todo{}
}
