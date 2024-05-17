package userservice

import (
	"errors"
	"halosuster/src/entities"
)

func (s *userService) Register(userRequest entities.User, isNurse bool) (entities.User, error) {
	if s.userRepository.FindByNIP(userRequest.NIP) {
		return entities.User{}, errors.New("NIP ALREADY EXIST")
	}

	return s.userRepository.Create(userRequest)
}
