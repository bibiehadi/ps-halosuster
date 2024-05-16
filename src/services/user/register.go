package userservice

import "halosuster/src/entities"

func (s *userService) Register(userRequest entities.NurseRequest) (entities.User, error) {
	user := entities.User{
		NIP:                 userRequest.NIP,
		Name:                userRequest.Name,
		IdentityCardScanImg: userRequest.IdentityCardScanImg,
		Role:                entities.Role(entities.Nurse),
		IsActive:            false,
	}

	return s.userRepository.Create(user)
}
