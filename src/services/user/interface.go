package userservice

import (
	"halosuster/src/entities"
	repositories "halosuster/src/repositories/user"
)

type UserService interface {
	Register(userRequest entities.User, isNurse bool) (entities.User, error)
	Update(userId string, nurseUpdateRequest entities.NurseUpdateRequest) error
	Delete(userId string) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func New(repository repositories.UserRepository) *userService {
	return &userService{repository}
}
