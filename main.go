package main

import (
	"context"
	userRest "dev-knowledge/adapters/controllers/rest"
	userUseCase "dev-knowledge/domain/useCase"
	"dev-knowledge/infrastructure/logger"
	"dev-knowledge/infrastructure/restServer"
	restServerController "dev-knowledge/infrastructure/restServer/controller"
	"dev-knowledge/infrastructure/restServer/response"
	init_services "dev-knowledge/init-services"
	"fmt"
)

const ServicePort = ":8080"

func main() {
	lightLogger := logger.NewLightLogger()
	server := restServer.NewFiberServer(lightLogger)
	useCase := userUseCase.NewUserUseCase()
	errRespService, err := response.NewErrorResponseService(response.NewErrorResolver(), lightLogger)
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
