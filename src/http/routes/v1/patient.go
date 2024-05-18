package v1

import (
	patientController "halosuster/src/http/controllers/patient"
	patientrepository "halosuster/src/repositories/patient"
	patientservice "halosuster/src/services/patient"
)

func (i *V1Routes) MountPatient() {
	g := i.Echo.Group("/patient")

	patientRepository := patientrepository.New(i.Db)
	patientService := patientservice.New(patientRepository)
	patientController := patientController.New(patientService)

	g.POST("/patient/register", patientController.Create)
}
