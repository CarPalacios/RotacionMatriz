package controllers

import (
	//logger "matriz/cmd/utils/log"
	//"matriz/pkg/services"
	//"matriz/pkg/entities"

	//"go.uber.org/zap"

	"github.com/gofiber/fiber/v2"
)

// Definición de las funciones controller

/**
func GetAll(c *fiber.Ctx) error {
	logger.Debug(c, "Inicio de método GetAll")
	list := services.GetAll(c)
	return c.Status(fiber.StatusOK).JSON(list)
}
*/

func RotateMatriz(matriz [][]int) [][]int {
	n := len(matriz)
    newMatriz := make([][]int, n)
    for i := range newMatriz {
        newMatriz[i] = make([]int, n)
    }
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            newMatriz[j][n-i-1] = matriz[i][j]
        }
    }
    return newMatriz
}

func HandleMatrizRotation(c *fiber.Ctx) error {
	var matriz [][]int

	if err := c.BodyParser(&matriz); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Formato de entrada no válido",
        })
    }

	rotatedMatriz := RotateMatriz(matriz)

	return c.JSON(rotatedMatriz)
}