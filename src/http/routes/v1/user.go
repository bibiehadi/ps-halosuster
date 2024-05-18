package v1

import (
	nurseController "halosuster/src/http/controllers/nurse"
	itController "halosuster/src/http/controllers/stafIT"
	"halosuster/src/http/middlewares"
	userrepository "halosuster/src/repositories/user"
	userservice "halosuster/src/services/user"
)

func (i *V1Routes) MountUser() {
	g := i.Echo.Group("/user")

	userRepository := userrepository.New(i.Db)
	userService := userservice.New(userRepository)
	nurseController := nurseController.New(userService)
	itController := itController.New(userService)

	g.GET("", nurseController.GetAll, middlewares.RequireAuth())

	g.POST("/it/register", itController.Register)
	g.POST("/it/login", itController.Login)

	g.POST("/nurse/register", nurseController.Register, middlewares.RequireAuth())
	g.POST("/nurse/login", nurseController.Login)
	g.PUT("/nurse/:id", nurseController.Update, middlewares.RequireAuth())
	g.DELETE("/nurse/:id", nurseController.Delete, middlewares.RequireAuth())
	g.POST("/nurse/:id/access", nurseController.Activate, middlewares.RequireAuth())
}
