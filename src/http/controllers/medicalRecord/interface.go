package v1medicalrecord

import (
	"github.com/go-playground/validator/v10"
	services "halosuster/src/services/medicalRecord"
)

type MedicalRecordController struct {
	medicalRecordService services.MedicalRecordService
	validator            *validator.Validate
}

func New(services services.MedicalRecordService) *MedicalRecordController {
	validate := validator.New()
	return &MedicalRecordController{services, validate}
}
