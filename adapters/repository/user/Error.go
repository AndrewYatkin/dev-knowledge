package userRepo

import (
	"context"
	userEntity "dev-knowledge/domain/entity/user"
	"dev-knowledge/infrastructure/errors"
	loggerInterface "dev-knowledge/infrastructure/logger/interface"
	"fmt"
)

var (
	ErrMongoRepoIsRequired    = errors.NewError("SYS", "MongoRepository is required")
	ErrLogPublisherIsRequired = errors.NewError("SYS", "LogPublished is required")
	ErrCollectionIsRequired   = errors.NewError("SYS", "Collection is required")

	ErrUserNotFound = errors.NewErrorWithLevel("feab43fa-001", "User not found", errors.Levels.Info())
	ErrInsertUser   = errors.NewErrorWithLevel("feab43fa-002", "Can not insert user", errors.Levels.Info())
	ErrUpdateUser   = errors.NewErrorWithLevel("feab43fa-004", "Can not update user", errors.Levels.Info())
	ErrDeleteUser   = errors.NewErrorWithLevel("feab43fa-005", "Can not delete user", errors.Levels.Info())
	ErrFindUser     = errors.NewErrorWithLevel("feab43fa-006", "Can not find user", errors.Levels.Info())
	ErrRestoreUser  = errors.NewErrorWithLevel("feab43fa-007", "Can not restore user", errors.Levels.Info())
)

type errorProcessor struct {
	logPublisher loggerInterface.LogPublisher
}

func (p *errorProcessor) LogAndReturnErrUserNotFoundByID(ctx context.Context, userID *userEntity.UserID) error {
	detailErrMsg := fmt.Sprintf("User by id = '%s' not found", userID)
	detailError := errors.NewError(ErrUserNotFound.Code(), detailErrMsg)
	p.logPublisher.LogError(ctx, detailError)

	return ErrUserNotFound
}

func (p *errorProcessor) LogAndReturnErrInsertUser(ctx context.Context, cause error) error {
	detailErrMsg := fmt.Sprintf("Fail insert user into database. Cause: '%s'", cause.Error())
	detailError := errors.NewError(ErrInsertUser.Code(), detailErrMsg)
	p.logPublisher.LogError(ctx, detailError)

	return ErrInsertUser
}

func (p *errorProcessor) LogAndReturnErrUpdateUser(ctx context.Context, userID *userEntity.UserID, cause error) error {
	detailErrMsg := fmt.Sprintf("Fail update user into database. UserID = '%s'. Cause: '%s'", userID, cause.Error())
	detailError := errors.NewError(ErrUpdateUser.Code(), detailErrMsg)
	p.logPublisher.LogError(ctx, detailError)

	return ErrUpdateUser
}

func (p *errorProcessor) LogAndReturnErrDeleteUser(ctx context.Context, cause error) error {
	detailErrMsg := fmt.Sprintf("Fail delete user from database. Cause: '%s'", cause.Error())
	detailError := errors.NewError(ErrDeleteUser.Code(), detailErrMsg)
	p.logPublisher.LogError(ctx, detailError)

	return ErrDeleteUser
}

func (p *errorProcessor) LogAndReturnErrFindUser(ctx context.Context, cause error) error {
	detailErrMsg := fmt.Sprintf("Fail find user into database. Cause: '%s'", cause.Error())
	detailError := errors.NewError(ErrFindUser.Code(), detailErrMsg)
	p.logPublisher.LogError(ctx, detailError)

	return ErrFindUser
}

func (p *errorProcessor) LogAndReturnErrRestoreUser(ctx context.Context, userID string, cause error) error {
	detailErrMsg := fmt.Sprintf("Fail restore user from database. UserID = '%s'. Cause: '%s'", userID, cause.Error())
	detailError := errors.NewError(ErrRestoreUser.Code(), detailErrMsg)
	p.logPublisher.LogError(ctx, detailError)

	return ErrRestoreUser
}
