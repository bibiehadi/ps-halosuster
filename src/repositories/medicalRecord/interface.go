package medicalrecordrepository

import (
	"halosuster/src/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MedicalRecordRepository interface {
	GetAllMedicalRecord(params entities.MedicalRecordQueryParams) ([]entities.MedicalRecordResponse, error)
	//GetMedicalRecordById(id int) (entities.MedicalRecord, error)
	CreateMedicalRecord(medicalRecord entities.MedicalRecordRequest) (entities.MedicalRecord, error)
}

type medicalRecordRepository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *medicalRecordRepository {
	return &medicalRecordRepository{db}
}
