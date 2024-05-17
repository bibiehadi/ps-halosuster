package userservice

import "halosuster/src/entities"

func (s *userService) Register(userRequest entities.User, isNurse bool) (entities.User, error) {
	return s.userRepository.Create(userRequest)
}
