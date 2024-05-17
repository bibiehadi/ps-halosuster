package userrepository

import (
	"halosuster/src/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Create(user entities.User) (entities.User, error)
	FindById(userId string) (entities.User, error)
	FindByNIP(nip int) bool
	Update(userId string, user entities.User) error
	Delete(userId string) error
	Activate(userId string, password string) error
}

type userRepository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *userRepository {
	return &userRepository{db}
}
