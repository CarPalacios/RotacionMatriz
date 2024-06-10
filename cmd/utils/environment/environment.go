package environment

import (
	"errors"
	"os"

	"matriz/cmd/utils/log"
)

var logger = log.GetLogger()

var Var1 string
var Var2 string

func ConfigVariables() error {
	Var1 = os.Getenv("VAR1")
	Var2 = os.Getenv("VAR2")

	// Verificar si alguna variable de entorno no está definida
	if !validateVariables() {
		msg := "faltan definiciones de variables de entorno"
		logger.Error(msg)
		return errors.New(msg)
	}
	return nil
}

func validateVariables() bool {
	// Descomentar y validar las variables necesarias
	/*if Var1 == "" || Var2 == "" {
		return false
	}*/
	return true
}
