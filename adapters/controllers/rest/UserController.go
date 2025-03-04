package userRest

import (
	userRestRequest "dev-knowledge/adapters/controllers/rest/request"
	"dev-knowledge/adapters/controllers/rest/serializer"
	usecaseInterface "dev-knowledge/boundary/domain/usecase"
	loggerInterface "dev-knowledge/infrastructure/logger/interface"
	restServerController "dev-knowledge/infrastructure/restServer/controller"
	"net/http"
)

type UserController struct {
	*restServerController.BaseController
	userUseCase usecaseInterface.UserUseCaseInterface
	logger      loggerInterface.Logger
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	requestData := &userRestRequest.CreateUserRequest{}
	if err := c.FillReqModel(r, requestData); err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	createdUser, err := c.userUseCase.Create(r.Context(), requestData.GetCreateUserDto())
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	response, err := serializer.SerializeUser(createdUser)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.JSONResponse(w, r, response, http.StatusCreated)
}
