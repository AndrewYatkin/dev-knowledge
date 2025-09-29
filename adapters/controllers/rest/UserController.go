package userRest

import (
	userRestRequest "dev-knowledge/adapters/controllers/rest/request"
	"dev-knowledge/adapters/controllers/rest/serializer"
	usecaseInterface "dev-knowledge/boundary/domain/usecase"
	"dev-knowledge/infrastructure/jwtService"
	logInterface "dev-knowledge/infrastructure/logger/interface"
	restServerController "dev-knowledge/infrastructure/restServer/controller"
	"net/http"
)

type UserController struct {
	*restServerController.BaseController
	userUseCase usecaseInterface.UserUseCaseInterface
	logger      logInterface.LogPublisher
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

func (c *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	executorRole, err := c.GetStrParamFromCtx(r.Context(), jwtService.UserRoleKey)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	userID, err := c.GetRouteParamFromCtx(r.Context(), "userID")
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	createdUser, err := c.userUseCase.GetUserByID(r.Context(), userID, executorRole)
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

func (c *UserController) Me(w http.ResponseWriter, r *http.Request) {
	executorUserID, err := c.GetStrParamFromCtx(r.Context(), jwtService.UserIDKey)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	user, err := c.userUseCase.GetMe(r.Context(), executorUserID)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	response, err := serializer.SerializeUser(user)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.JSONResponse(w, r, response, http.StatusOK)
}
