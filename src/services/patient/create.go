package patientservice

import (
	"errors"
	"halosuster/src/entities"
)

func (s *patientService) Create(patientRequest entities.Patient) (entities.Patient, error) {
	if s.patientRepository.IDisExist(patientRequest.IdentityNumber) {
		return entities.Patient{}, errors.New("IDENTITY NUMBER ALREADY EXIST")
	}

	return s.patientRepository.Create(patientRequest)
}
