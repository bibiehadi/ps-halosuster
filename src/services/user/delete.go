package userservice

import "errors"

func (s *userService) Delete(userId string) error {
	user, findError := s.userRepository.FindById(userId)

	if findError != nil {
		return findError
	}

	if user.Role != "nurse" {
		return errors.New("THIS USER IS NOT NURSE")
	}

	err := s.userRepository.Delete(userId)
	if err != nil {
		return err
	}

	return nil
}
