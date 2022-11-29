package repository

import "todo/model"

//go:generate mockgen -source=todo_repository.go -destination=../mock/mock_repository.go -package=mock
type ITodoRepository interface {
	CreateTodo(todoObj model.Todo) model.Todo
}

type TodoRepository struct {
}

func NewTodoRepository() ITodoRepository {
	return &TodoRepository{}
}

func (r *TodoRepository) CreateTodo(todoObj model.Todo) model.Todo {
	return todoObj
}
