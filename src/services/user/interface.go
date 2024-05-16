package userservice

import (
	"halosuster/src/entities"
	repositories "halosuster/src/repositories/user"
)

type UserService interface {
	Register(userRequest entities.NurseRequest) (entities.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func New(repository repositories.UserRepository) *userService {
	return &userService{repository}
}
