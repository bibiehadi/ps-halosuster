package userservice

import (
	"errors"
	"halosuster/src/entities"
	"time"
)

func (s *userService) Update(userId string, nurseUpdateRequest entities.NurseUpdateRequest) error {
	user, findError := s.userRepository.FindById(userId)

	if findError != nil {
		return findError
	}

	if s.userRepository.NIPisExist(nurseUpdateRequest.NIP) && user.NIP != nurseUpdateRequest.NIP {
		return errors.New("NIP ALREADY EXIST")
	}

	if user.Role != "nurse" {
		return errors.New("THIS USER IS NOT NURSE")
	}

	nurse := entities.User{
		ID:                  user.ID,
		NIP:                 nurseUpdateRequest.NIP,
		Name:                nurseUpdateRequest.Name,
		IdentityCardScanImg: user.IdentityCardScanImg,
		Role:                entities.Nurse,
		IsActive:            user.IsActive,
		Password:            user.Password,
		CreatedAt:           user.CreatedAt,
		UpdatedAt:           time.Now(),
	}

	err := s.userRepository.Update(userId, nurse)
	if err != nil {
		return err
	}

	return nil
}
