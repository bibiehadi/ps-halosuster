package userservice

import (
	"errors"
	"fmt"
	"halosuster/src/entities"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (s *userService) Login(authRequest entities.AuthRequest) (string, entities.User, error) {
	user, err := s.userRepository.FindByNIP(authRequest.NIP)

	if err != nil {
		fmt.Println("invalid NIP")
		return "", entities.User{}, errors.New("INVALID NIP OR PASSWORD")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authRequest.Password))
	if err != nil {
		fmt.Println("failed matching")
		return "", entities.User{}, errors.New("INVALID NIP OR PASSWORD")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.NIP,
		"exp": time.Now().Add(time.Hour * 8).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		fmt.Println("error signed string")
		return "", entities.User{}, err
	}

	return tokenString, user, err
}
