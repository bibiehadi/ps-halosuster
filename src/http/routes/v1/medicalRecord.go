package v1

import (
	medicalRecordController "halosuster/src/http/controllers/medicalRecord"
	"halosuster/src/http/middlewares"
	medicalRecordrepository "halosuster/src/repositories/medicalRecord"
	medicalRecordservice "halosuster/src/services/medicalRecord"
)

func (i *V1Routes) MountMedicalRecords() {
	g := i.Echo.Group("/medical")
	g.Use(middlewares.RequireAuth())
	medicalRecordRepository := medicalRecordrepository.New(i.Db)
	medicalRecordService := medicalRecordservice.New(medicalRecordRepository)
	medicalRecordController := medicalRecordController.New(medicalRecordService)

	g.POST("/record", medicalRecordController.CreateMedicalRecord)
	g.GET("/record", medicalRecordController.GetAll)
}
