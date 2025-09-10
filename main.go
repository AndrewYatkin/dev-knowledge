package main

import (
	"context"
	userRepo "dev-knowledge/adapters/controllers/repository"
	userRest "dev-knowledge/adapters/controllers/rest"
	restResolver "dev-knowledge/adapters/controllers/rest/resolver"
	userUseCase "dev-knowledge/domain/useCase"
	jwtservice "dev-knowledge/infrastructure/jwtService"
	"dev-knowledge/infrastructure/logger"
	"dev-knowledge/infrastructure/restServer"
	restServerController "dev-knowledge/infrastructure/restServer/controller"
	"dev-knowledge/infrastructure/restServer/response"
	init_services "dev-knowledge/init-services"
	"fmt"
)

const ServicePort = ":8081"

func main() {
	lightLogger := logger.NewLightLogger()
	jwtService, err := jwtservice.NewBuilder().Secret("1").Build()
	if err != nil {
		stopService(lightLogger, err)
		return
	}
	server := restServer.NewFiberServer(lightLogger, jwtService)
	userRepoDummy, err := userRepo.NewBuilder().Logger(lightLogger).Build()
	if err != nil {
		stopService(lightLogger, err)
		return
	}
	useCase, err := userUseCase.NewBuilder().
		JwtService(jwtService).
		UserRepo(userRepoDummy).
		Build()
	if err != nil {
		stopService(lightLogger, err)
		return
	}
	errRespService, err := response.NewErrorResponseService(restResolver.NewErrorResolver(), lightLogger)
	if err != nil {
		stopService(lightLogger, err)
		return
	}
	responseService, err := response.NewResponseService(errRespService, lightLogger)
	if err != nil {
		stopService(lightLogger, err)
		return
	}
	baseController, err := restServerController.NewBaseController(responseService, lightLogger)
	if err != nil {
		stopService(lightLogger, err)
		return
	}
	userController, err := userRest.NewBuilder().
		BaseController(baseController).
		UserUseCase(useCase).
		Logger(lightLogger).
		Build()
	if err != nil {
		stopService(lightLogger, err)
		return
	}
	router := init_services.NewUserRouter(server, userController)
	router.RegisterRoutes()

	lightLogger.Info(context.Background(), "server is starting...")
	if err := server.Start(ServicePort); err != nil {
		lightLogger.Error(context.Background(), fmt.Errorf("failed to start server: %v", err))
	}
}

func stopService(lightLogger *logger.LightLogger, err error) {
	lightLogger.Error(context.Background(), err)
	lightLogger.Info(context.Background(), " stopped")
}
