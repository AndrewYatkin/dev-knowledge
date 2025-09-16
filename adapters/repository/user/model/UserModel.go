package userRepoModel

import (
	userEntity "dev-knowledge/domain/entity/user"
	emailEntity "dev-knowledge/domain/entity/user/email"
	passwordEntity "dev-knowledge/domain/entity/user/password"
	"dev-knowledge/domain/entity/user/spec"
	commonTime "dev-knowledge/infrastructure/tools/time"
)

type User struct {
	ID        string     `bson:"user_id"`
	Profile   *Profile   `bson:"profile"`
	Role      string     `bson:"role"`
	Password  *Password  `bson:"password"`
	Email     *Emails    `bson:"email"`
	Agreement *Agreement `bson:"agreement"`
	CreatedAt int64      `bson:"created_at"`
	Removed   bool       `bson:"removed"`
}

func UserToEntity(repoUser *User) (*userEntity.User, error) {
	userID, err := userEntity.UserIDFrom(repoUser.ID)
	if err != nil {
		return nil, err
	}

	var emails *emailEntity.UserEmails
	if repoUser.Email != nil {
		emails, err = EmailsToEntity(repoUser.Email)
		if err != nil {
			return nil, err
		}
	}

	var password *passwordEntity.Password
	if repoUser.Password != nil {
		password, err = PasswordToEntity(repoUser.Password)
		if err != nil {
			return nil, err
		}
	}
	userCreatedAt := commonTime.FromUnixNano(repoUser.CreatedAt)

	agreement, err := AgreementToEntity(repoUser.Agreement)
	if err != nil {
		return nil, err
	}

	profile, err := ProfileToEntity(repoUser.Profile)
	if err != nil {
		return nil, err
	}

	role, err := spec.UserRoles.Of(repoUser.Role)
	if err != nil {
		return nil, err
	}

	userBuild := userEntity.NewBuilder().
		ID(userID).
		Email(emails).
		Profile(profile).
		Role(role).
		Agreement(agreement).
		Password(password).
		CreatedAt(userCreatedAt).
		Removed(repoUser.Removed)

	return userBuild.Build()
}

func UserToModel(user *userEntity.User) *User {

	var userEmailsModel *Emails
	if userEmailsEntity, exists := user.Emails(); exists {
		userEmailsModel = EmailsToModel(userEmailsEntity)
	}

	var agreementModel *Agreement
	if agreement, ok := user.Agreement(); ok {
		agreementModel = AgreementToModel(&agreement)
	}

	var passwordModel *Password
	if password, ok := user.Password(); ok {
		passwordModel = PasswordToModel(&password)
	}

	userModel := &User{
		ID:        user.ID().String(),
		Profile:   getUserProfileModel(user),
		Role:      user.Role().String(),
		Password:  passwordModel,
		Email:     userEmailsModel,
		Agreement: agreementModel,
		CreatedAt: user.CreatedAt().UnixNano(),
		Removed:   user.Removed(),
	}

	return userModel
}

func getUserProfileModel(user *userEntity.User) *Profile {
	profile := user.Profile()
	return ProfileToModel(&profile)
}
