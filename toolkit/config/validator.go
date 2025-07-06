package config

import (
	"github.com/go-playground/validator/v10"

	validatorHandler "abarobotics-test/src/handler/validator"
)

// NewValidator.
func NewValidator() *validatorHandler.Validator {
	return &validatorHandler.Validator{Validator: validator.New()}
}
