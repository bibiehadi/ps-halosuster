package v1

import (
	nurseController "halosuster/src/http/controllers/nurse"
	itController "halosuster/src/http/controllers/stafIT"
	userrepository "halosuster/src/repositories/user"
	userservice "halosuster/src/services/user"
)

func (i *V1Routes) MountUser() {
	g := i.Echo.Group("/user")

	userRepository := userrepository.New(i.Db)
	userService := userservice.New(userRepository)
	nurseController := nurseController.New(userService)
	itController := itController.New(userService)

	g.GET("", nurseController.GetAll)

	g.POST("/it/register", itController.Register)
	g.POST("/it/login", itController.Login)

	g.POST("/nurse/register", nurseController.Register)
	g.POST("/nurse/login", nurseController.Login)
	g.PUT("/nurse/:id", nurseController.Update)
	g.DELETE("/nurse/:id", nurseController.Delete)
	g.POST("/nurse/:id/access", nurseController.Activate)
}
