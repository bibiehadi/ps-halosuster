package userservice

import (
	"halosuster/src/entities"

	"golang.org/x/crypto/bcrypt"
)

func (s *userService) Activate(userId string, activateRequest entities.NurseActivate) error {
	_, findError := s.userRepository.FindById(userId)
	if findError != nil {
		return findError
	}

	hashed, hashErr := bcrypt.GenerateFromPassword([]byte(activateRequest.Password), 10)
	if hashErr != nil {
		return hashErr
	}

	err := s.userRepository.Activate(userId, string(hashed))
	if err != nil {
		return err
	}

	return nil
}
