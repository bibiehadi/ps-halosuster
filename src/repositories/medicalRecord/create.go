package medicalrecordrepository

import (
	"context"
	"halosuster/src/entities"
)

func (r *medicalRecordRepository) CreateMedicalRecord(medicalRecordRequest entities.MedicalRecordRequest) (entities.MedicalRecord, error) {
	var query string = "INSERT INTO medical_records (pasient_id, sympthoms, medications, created_by) values ($1,$2,$3,$4) RETURNING id"
	var medicalRecordId int
	err := r.db.QueryRow(context.Background(), query, medicalRecordRequest.IdentityNumber, medicalRecordRequest.Symptoms, medicalRecordRequest.Medications, medicalRecordRequest.CreatedBy).Scan(
		&medicalRecordId,
	)
	if err != nil {
		return entities.MedicalRecord{}, err
	}
	medicalRecord := entities.MedicalRecord{
		ID: medicalRecordId,
	}
	return medicalRecord, err
}
