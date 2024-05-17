package userservice

func (s *userService) Delete(userId string) error {
	_, findError := s.userRepository.FindById(userId)

	if findError != nil {
		return findError
	}

	err := s.userRepository.Delete(userId)
	if err != nil {
		return err
	}

	return nil
}
