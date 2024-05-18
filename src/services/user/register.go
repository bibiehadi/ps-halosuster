package userservice

import (
	"errors"
	"halosuster/src/entities"
	"halosuster/src/helpers"
)

func (s *userService) Register(userRequest entities.User, isNurse bool) (entities.User, error) {
	if !isNurse {
		if s.userRepository.NIPisExist(userRequest.NIP) {
			return entities.User{}, errors.New("NIP ALREADY EXIST")
		}

		hashPassword, hashErr := helpers.HashPassword(userRequest.Password)

		if hashErr != nil {
			return entities.User{}, hashErr
		}
		stafIT := entities.User{
			NIP:      userRequest.NIP,
			Name:     userRequest.Name,
			Password: hashPassword,
			Role:     entities.IT,
			IsActive: true,
		}
		return s.userRepository.Create(stafIT)
	}

	if s.userRepository.NIPisExist(userRequest.NIP) {
		return entities.User{}, errors.New("NIP ALREADY EXIST")
	}

	return s.userRepository.Create(userRequest)
}
