package initServices

import (
	userRest "dev-knowledge/adapters/controllers/rest"
	restServerInterface "dev-knowledge/infrastructure/restServer/interface"
	"net/http"
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
	r.server.RegisterPublicRoute(http.MethodPost, "/user/create", r.controller.CreateUser)
	r.server.RegisterPrivateRoute(http.MethodGet, "/user/me", r.controller.Me)
	r.server.RegisterPrivateRoute(http.MethodGet, "/user/:userID", r.controller.GetUserById)
}
