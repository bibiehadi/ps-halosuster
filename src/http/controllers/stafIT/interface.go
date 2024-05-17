package itcontroller

import (
	services "halosuster/src/services/user"

	"github.com/go-playground/validator/v10"
)

type stafItController struct {
	userService services.UserService
	validator   *validator.Validate
}

func New(services services.UserService) *stafItController {
	validate := validator.New()
	return &stafItController{services, validate}
}
