package medicalrecordservice

import "halosuster/src/entities"

func (s *medicalRecordService) CreateMedicalRecord(medicalRecordRequest entities.MedicalRecordRequest) (entities.MedicalRecordResponse, error) {
	medicalRecord, err := s.medicalRecordRepository.CreateMedicalRecord(medicalRecordRequest)
	if err != nil {
		return entities.MedicalRecordResponse{}, err
	}
	return medicalRecord, nil

}
