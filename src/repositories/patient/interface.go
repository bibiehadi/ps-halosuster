package patientrepository

import (
	"halosuster/src/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PatientRepository interface {
	Create(user entities.Patient) (entities.Patient, error)
	FindAll(params entities.PatientQueryParams) ([]entities.PatientResponse, error)
	IDisExist(patientId int64) bool
}

type patientRepository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *patientRepository {
	return &patientRepository{db}
}
