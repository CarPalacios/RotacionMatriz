package routes

import (
	//"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"matriz/api/controllers"
)

func GoFiberRoutes(app *fiber.App) {
	// Definicion de los routes
	app.Post("/rotate", controllers.HandleMatrizRotation)
}
