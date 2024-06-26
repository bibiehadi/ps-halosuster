package userrepository

import (
	"halosuster/src/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Create(user entities.User) (entities.User, error)
	GetAll(params entities.UserQueryParams) ([]entities.UserResponse, error)
	FindById(userId string) (entities.User, error)
	NIPisExist(nip int) bool
	FindByNIP(nip int) (entities.User, error)
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
