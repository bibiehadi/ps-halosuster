package v1

import (
	nurseController "halosuster/src/http/controllers/nurse"
	userrepository "halosuster/src/repositories/user"
	userservice "halosuster/src/services/user"
)

func (i *V1Routes) MountUser() {
	g := i.Echo.Group("/user")

	userRepository := userrepository.New(i.Db)
	userService := userservice.New(userRepository)
	nurseController := nurseController.New(userService)

	g.GET("", nurseController.GetAll)
	g.POST("/nurse/register", nurseController.Register)
	g.PUT("/nurse/:id", nurseController.Update)
	g.DELETE("/nurse/:id", nurseController.Delete)
	g.POST("/nurse/:id/access", nurseController.Activate)
}
