package main

import (
	"log"
	"matriz/api/routes"
	//"matriz/cmd/server"
	//"matriz/cmd/utils/environment"

	"github.com/gofiber/fiber/v2"
	//"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	routes.GoFiberRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
