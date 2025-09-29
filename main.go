package main

import (
	"context"
	userRepo "dev-knowledge/adapters/controllers/repository"
	userRest "dev-knowledge/adapters/controllers/rest"
	restResolver "dev-knowledge/adapters/controllers/rest/resolver"
	userUseCase "dev-knowledge/domain/useCase"
	jwtservice "dev-knowledge/infrastructure/jwtService"
	logInterface "dev-knowledge/infrastructure/logger/interface"
	"dev-knowledge/infrastructure/restServer"
	restServerController "dev-knowledge/infrastructure/restServer/controller"
	"dev-knowledge/infrastructure/restServer/response"
	initServices "dev-knowledge/init-services"
	"fmt"
)

const ServicePort = ":8081"
const Environment = "develop"
const AppID = "example"
const JWTSecret = "1"

var StopChan = make(chan struct{})

func main() {
	initLogger := initServices.NewLoggerInit()
	logPublisher := initLogger.Init(AppID, Environment, StopChan)
	jwtService, err := jwtservice.NewBuilder().Secret(JWTSecret).Build()
	if err != nil {
		stopService(logPublisher, err)
		return
	}
	server := restServer.NewFiberServer(logPublisher, jwtService)
	userRepoDummy, err := userRepo.NewBuilder().Logger(logPublisher).Build()
	if err != nil {
		stopService(logPublisher, err)
		return
	}
	useCase, err := userUseCase.NewBuilder().
		JwtService(jwtService).
		UserRepo(userRepoDummy).
		Build()
	if err != nil {
		stopService(logPublisher, err)
		return
	}
	errRespService, err := response.NewErrorResponseService(restResolver.NewErrorResolver(), logPublisher)
	if err != nil {
		stopService(logPublisher, err)
		return
	}
	responseService, err := response.NewResponseService(errRespService, logPublisher)
	if err != nil {
		stopService(logPublisher, err)
		return
	}
	baseController, err := restServerController.NewBaseController(responseService, logPublisher)
	if err != nil {
		stopService(logPublisher, err)
		return
	}
	userController, err := userRest.NewBuilder().
		BaseController(baseController).
		UserUseCase(useCase).
		Logger(logPublisher).
		Build()
	if err != nil {
		stopService(logPublisher, err)
		return
	}
	router := initServices.NewUserRouter(server, userController)
	router.RegisterRoutes()

	logPublisher.LogInfo(context.Background(), "server is starting...")
	if err := server.Start(ServicePort); err != nil {
		logPublisher.LogError(context.Background(), fmt.Errorf("failed to start server: %v", err))
	}
}

func stopService(logPublisher logInterface.LogPublisher, err error) {
	logPublisher.LogError(context.Background(), err)
	logPublisher.LogInfo(context.Background(), " stopped")
}
