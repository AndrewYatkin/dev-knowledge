package repositoryInterface

import (
	"context"
	emailPrimitive "dev-knowledge/common/domainPrimitive/primitive/email"
	userEntity "dev-knowledge/domain/entity/user"
	emailEntity "dev-knowledge/domain/entity/user/email"
	passwordEntity "dev-knowledge/domain/entity/user/password"
	profileEntity "dev-knowledge/domain/entity/user/profile"
)

type UserRepository interface {
	Insert(ctx context.Context, user *userEntity.User) error
	Update(ctx context.Context, user *userEntity.User) error
	UpdatePassword(ctx context.Context, userID *userEntity.UserID, password *passwordEntity.Password) error
	UpdateProfile(ctx context.Context, userID *userEntity.UserID, profile *profileEntity.Profile) error
	UpdateEmails(ctx context.Context, userID *userEntity.UserID, email *emailEntity.UserEmails) error

	GetUserByID(ctx context.Context, userID *userEntity.UserID) (*userEntity.User, error)
	GetUserByActivatedEmail(ctx context.Context, email emailPrimitive.Email) (*userEntity.User, error)
	RemoveByID(ctx context.Context, userID *userEntity.UserID) error
}
