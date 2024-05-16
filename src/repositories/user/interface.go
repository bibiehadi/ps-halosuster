package userrepository

import (
	"halosuster/src/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Create(user entities.User) (entities.User, error)
}

type userRepository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *userRepository {
	return &userRepository{db}
}
