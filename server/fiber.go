package server

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"todo/handler"
	"todo/repository"
	"todo/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Router() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}

	err = StartServer(port)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Provider Service Listening :3000")
}

func StartServer(port int) error {

	app := fiber.New()

	app.Use(cors.New())

	handler := handler.NewTodoHandler(
		service.NewTodoService(
			repository.NewTodoRepository(),
		),
	)

	app.Post("/todos", handler.CreateTodo)

	return app.Listen(fmt.Sprintf(":%d", port))
}
