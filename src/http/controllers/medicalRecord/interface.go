package v1medicalrecord

import (
	"github.com/go-playground/validator/v10"
	services "halosuster/src/services/medicalRecord"
	patientservice "halosuster/src/services/patient"
)

type MedicalRecordController struct {
	medicalRecordService services.MedicalRecordService
	patientService       patientservice.PatientService
	validator            *validator.Validate
}

func New(services services.MedicalRecordService, patientService patientservice.PatientService) *MedicalRecordController {
	validate := validator.New()
	return &MedicalRecordController{services, patientService, validate}
}
