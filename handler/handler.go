package handler

import (
	"github.com/gofiber/fiber/v2"
	"todo/service"
)

type ITodoHandler interface {
	CreateTodo(ctx *fiber.Ctx) error
}

type TodoHandler struct {
	service service.ITodoService
}

func NewTodoHandler(service service.ITodoService) ITodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) CreateTodo(ctx *fiber.Ctx) error {
	todo := ctx.FormValue("todo")
	return ctx.JSON(h.service.CreateTodo(todo))
}
