package v1

import (
	medicalRecordController "halosuster/src/http/controllers/medicalRecord"
	"halosuster/src/http/middlewares"
	medicalRecordrepository "halosuster/src/repositories/medicalRecord"
	patientrepository "halosuster/src/repositories/patient"
	medicalRecordservice "halosuster/src/services/medicalRecord"
	patientservice "halosuster/src/services/patient"
)

func (i *V1Routes) MountMedicalRecords() {
	g := i.Echo.Group("/medical")
	g.Use(middlewares.RequireAuth())
	medicalRecordRepository := medicalRecordrepository.New(i.Db)
	medicalRecordService := medicalRecordservice.New(medicalRecordRepository)
	patientRepository := patientrepository.New(i.Db)
	patientService := patientservice.New(patientRepository)
	medicalRecordController := medicalRecordController.New(medicalRecordService, patientService)

	g.POST("/record", medicalRecordController.CreateMedicalRecord)
	g.GET("/record", medicalRecordController.GetAll)
}
