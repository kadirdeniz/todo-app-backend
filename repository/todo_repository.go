package repository

import (
	"todo/model"
)

var todoList []model.Todo

//go:generate mockgen -source=todo_repository.go -destination=../mock/mock_repository.go -package=mock
type ITodoRepository interface {
	CreateTodo(todoList *[]model.Todo, todoObj model.Todo) model.Todo
}

type TodoRepository struct {
}

func NewTodoRepository() ITodoRepository {
	return &TodoRepository{}
}

func (r *TodoRepository) CreateTodo(todoList *[]model.Todo, todoObj model.Todo) model.Todo {
	*todoList = append(*todoList, todoObj)

	return todoObj
}
