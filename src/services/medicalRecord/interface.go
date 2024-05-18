package medicalrecordservice

import (
	"halosuster/src/entities"
	repositories "halosuster/src/repositories/medicalRecord"
)

type MedicalRecordService interface {
	CreateMedicalRecord(medicalRecordRequest entities.MedicalRecordRequest) (entities.MedicalRecordResponse, error)
	GetAll(params entities.MedicalRecordQueryParams) ([]entities.MedicalRecordResponse, error)
}

type medicalRecordService struct {
	medicalRecordRepository repositories.MedicalRecordRepository
}

func New(repository repositories.MedicalRecordRepository) *medicalRecordService {
	return &medicalRecordService{repository}
}
