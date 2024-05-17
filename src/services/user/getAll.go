package userservice

import "halosuster/src/entities"

func (s *userService) GetAll(params entities.UserQueryParams) ([]entities.UserResponse, error) {
	return s.userRepository.GetAll(params)
}
