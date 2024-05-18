package medicalrecordrepository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"halosuster/src/entities"
)

type MedicalRecordRepository interface {
	GetAllMedicalRecord() ([]entities.MedicalRecord, error)
	GetMedicalRecordById(id int) (entities.MedicalRecord, error)
	CreateMedicalRecord(medicalRecord entities.MedicalRecordRequest) (entities.MedicalRecord, error)
}

type medicalRecordRepository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *medicalRecordRepository {
	return &medicalRecordRepository{db}
}
