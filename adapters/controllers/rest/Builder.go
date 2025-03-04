package userRest

import (
	usecaseInterface "dev-knowledge/boundary/domain/usecase"
	"dev-knowledge/infrastructure/errors"
	loggerInterface "dev-knowledge/infrastructure/logger/interface"
	restServerController "dev-knowledge/infrastructure/restServer/controller"
)

type UserControllerBuilder struct {
	controller *UserController
	errors     *errors.Errors
}

func NewBuilder() *UserControllerBuilder {
	return &UserControllerBuilder{
		controller: &UserController{},
		errors:     errors.NewErrors(),
	}
}

func (b *UserControllerBuilder) Logger(logger loggerInterface.Logger) *UserControllerBuilder {
	b.controller.logger = logger
	return b
}

func (b *UserControllerBuilder) BaseController(base *restServerController.BaseController) *UserControllerBuilder {
	b.controller.BaseController = base
	return b
}

func (b *UserControllerBuilder) UserUseCase(userUseCase usecaseInterface.UserUseCaseInterface) *UserControllerBuilder {
	b.controller.userUseCase = userUseCase
	return b
}

func (b *UserControllerBuilder) Build() (*UserController, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	b.fillDefaultFields()
	return b.controller, nil
}

func (b *UserControllerBuilder) checkRequiredFields() {
	if b.controller.BaseController == nil {
		b.errors.AddError(errors.NewError("SYS", "UserControllerBuilder: BaseController is required"))
	}
	if b.controller.userUseCase == nil {
		b.errors.AddError(errors.NewError("SYS", "UserControllerBuilder: UserUseCase is required"))
	}
}

func (b *UserControllerBuilder) fillDefaultFields() {}
