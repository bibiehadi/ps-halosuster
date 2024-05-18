package patientservice

import (
	"halosuster/src/entities"
)

func (s *patientService) Create(patientRequest entities.Patient) (entities.Patient, error) {
	return s.patientRepository.Create(patientRequest)
}
