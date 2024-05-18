package patientcontroller

import (
	services "halosuster/src/services/patient"

	"github.com/go-playground/validator/v10"
)

type patientController struct {
	patientService services.PatientService
	validator      *validator.Validate
}

func New(services services.PatientService) *patientController {
	validate := validator.New()
	return &patientController{services, validate}
}
