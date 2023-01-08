package server

import (
	"fmt"
	"log"
	"todo/handler"
	"todo/repository"
	"todo/service"

	"github.com/gofiber/fiber/v2"
)

func Router() {
	err := StartServer(8000)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Provider Service Listening :3000")
}

func StartServer(port int) error {

	app := fiber.New()

	handler := handler.NewTodoHandler(
		service.NewTodoService(
			repository.NewTodoRepository(),
		),
	)

	app.Post("/todos", handler.CreateTodo)

	return app.Listen(fmt.Sprintf(":%d", port))
}
