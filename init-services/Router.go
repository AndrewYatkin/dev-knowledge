package initServices

import (
	userRest "dev-knowledge/adapters/controllers/rest"
	restServerInterface "dev-knowledge/infrastructure/restServer/interface"
)

type Router interface {
	RegisterRoutes()
}

type UserRouter struct {
	server     restServerInterface.Server
	controller *userRest.UserController
}

func NewUserRouter(
	server restServerInterface.Server,
	controller *userRest.UserController,
) *UserRouter {
	return &UserRouter{
		controller: controller,
		server:     server,
	}
}

func (r *UserRouter) RegisterRoutes() {
	r.server.RegisterPublicRoute("POST", "/user/create", r.controller.CreateUser)
}
