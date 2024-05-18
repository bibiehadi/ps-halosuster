package medicalrecordservice

import "halosuster/src/entities"

func (s *medicalRecordService) CreateMedicalRecord(medicalRecordRequest entities.MedicalRecordRequest) (entities.MedicalRecord, error) {
	medicalRecord, err := s.medicalRecordRepository.CreateMedicalRecord(medicalRecordRequest)
	if err != nil {
		return entities.MedicalRecord{}, err
	}
	return medicalRecord, nil

}
