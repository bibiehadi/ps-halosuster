package userservice

import (
	"halosuster/src/entities"
	repositories "halosuster/src/repositories/user"
)

type UserService interface {
	Register(userRequest entities.User, isNurse bool) (entities.User, error)
	GetAll(params entities.UserQueryParams) ([]entities.UserResponse, error)
	Update(userId string, nurseUpdateRequest entities.NurseUpdateRequest) error
	Delete(userId string) error
	Activate(userId string, activateRequest entities.NurseActivate) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func New(repository repositories.UserRepository) *userService {
	return &userService{repository}
}
