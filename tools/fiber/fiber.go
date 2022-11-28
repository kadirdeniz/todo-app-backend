package fiber

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func Router() {
	err := StartServer(8000)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Provider Service Listening :3000")
}

func StartServer(port int) error {

	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	return app.Listen(fmt.Sprintf(":%d", port))
}
