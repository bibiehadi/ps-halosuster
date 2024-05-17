package userservice

import (
	"errors"
	"halosuster/src/entities"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func (s *userService) Register(userRequest entities.User, isNurse bool) (entities.User, error) {
	if !isNurse {
		if s.userRepository.NIPisExist(userRequest.NIP) {
			return entities.User{}, errors.New("NIP ALREADY EXIST")
		}
		salt, _ := strconv.Atoi(os.Getenv("BCRYPT_SALT"))
		hashed, hashErr := bcrypt.GenerateFromPassword([]byte(userRequest.Password), salt)
		if hashErr != nil {
			return entities.User{}, hashErr
		}
		stafIT := entities.User{
			NIP:      userRequest.NIP,
			Name:     userRequest.Name,
			Password: string(hashed),
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
