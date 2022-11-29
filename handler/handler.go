package handler

import (
	"github.com/gofiber/fiber/v2"
	"todo/service"
)

type CreatTodoRequest struct {
	Text string `json:"text"`
}

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
	request := CreatTodoRequest{}
	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	return ctx.Status(200).JSON(h.service.CreateTodo(request.Text))
}
