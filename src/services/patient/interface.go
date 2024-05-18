package patientservice

import (
	"halosuster/src/entities"
	repositories "halosuster/src/repositories/patient"
)

type PatientService interface {
	IDisExist(patientId int64) bool
	Create(patient entities.Patient) (entities.Patient, error)
	// GetAll(params entities.PatientQueryParams) ([]entities.PatientResponse, error)
	// Update(patientId string, nurseUpdateRequest entities.NurseUpdateRequest) error
}

type patientService struct {
	patientRepository repositories.PatientRepository
}

func New(repository repositories.PatientRepository) *patientService {
	return &patientService{repository}
}
