package handler

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type UserRequest struct {
	Username string `validate:"required,min=1,max=50"`
	Password string `validate:"required,min=1,max=255"`
}
