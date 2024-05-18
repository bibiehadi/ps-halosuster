package medicalrecordservice

import "halosuster/src/entities"

func (s *medicalRecordService) GetAll(params entities.MedicalRecordQueryParams) ([]entities.MedicalRecordResponse, error) {
	medicalRecords, err := s.medicalRecordRepository.GetAllMedicalRecord(params)
	if err != nil {
		return []entities.MedicalRecordResponse{}, err
	}
	return medicalRecords, nil

}
