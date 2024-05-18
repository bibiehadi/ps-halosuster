package patientrepository

import (
	"halosuster/src/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PatientRepository interface {
	Create(user entities.Patient) (entities.Patient, error)
	// GetAll(params entities.PatientQueryParams) ([]entities.PatientResponse, error)
	// FindById(userId string) (entities.User, error)
	IDisExist(patientId int64) bool
	// Update(userId string, user entities.User) error
}

type patientRepository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *patientRepository {
	return &patientRepository{db}
}
