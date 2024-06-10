package services

import (
	logger "matriz/cmd/utils/log"
	"matriz/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

// Implementacion de funciones
func GetAll(c *fiber.Ctx) []entities.MatrizEntity {
	logger.Debug(c, "Inicio del metodo GetAll")
	return nil
}
