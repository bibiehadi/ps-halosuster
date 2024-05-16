package nursecontroller

import (
	services "halosuster/src/services/user"

	"github.com/go-playground/validator/v10"
)

type nurseController struct {
	userService services.UserService
	validator   *validator.Validate
}

func New(services services.UserService) *nurseController {
	validate := validator.New()
	return &nurseController{services, validate}
}
