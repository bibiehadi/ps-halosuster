package patientservice

import "halosuster/src/entities"

func (s *patientService) FindAll(params entities.PatientQueryParams) ([]entities.PatientResponse, error) {
	return s.patientRepository.FindAll(params)
}
