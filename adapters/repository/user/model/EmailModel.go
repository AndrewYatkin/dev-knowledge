package userRepoModel

import (
	emailPrimitive "dev-knowledge/common/domainPrimitive/primitive/email"
	verificationPrimitive "dev-knowledge/common/domainPrimitive/primitive/verification"
	emailEntity "dev-knowledge/domain/entity/user/email"
	commonTime "dev-knowledge/infrastructure/tools/time"
)

type Emails struct {
	ActivatedEmail    *UserEmail `bson:"activated_email"`
	NonActivatedEmail *UserEmail `bson:"non_activated_email"`
}

type UserEmail struct {
	ID               string            `bson:"user_id"`
	Email            string            `bson:"email"`
	VerificationCode *VerifivationCode `bson:"verification_code"`
	CreatedAt        int64             `bson:"created_at"`
}

func EmailsToEntity(repoEmails *Emails) (*emailEntity.UserEmails, error) {
	var err error
	var activatedEmail *emailEntity.UserEmail
	if repoEmails.ActivatedEmail != nil {
		activatedEmail, err = getEmailEntity(repoEmails.ActivatedEmail)
		if err != nil {
			return nil, err
		}
	}

	var notActivatedEmail *emailEntity.UserEmail
	if repoEmails.NotActivatedEmail != nil {
		notActivatedEmail, err = getEmailEntity(repoEmails.NotActivatedEmail)
		if err != nil {
			return nil, err
		}
	}

	return emailEntity.EmailsFrom(notActivatedEmail, activatedEmail), nil
}

func EmailsToModel(emails *emailEntity.UserEmails) *Emails {
	var activatedEmailEntity *emailEntity.UserEmail
	var activatedEmailModel *UserEmail
	if ap, ok := emails.ActivatedEmailData(); ok {
		activatedEmailEntity = &ap
		activatedEmailModel = EmailToModel(activatedEmailEntity)
	}

	var notActivatedEmailEntity *emailEntity.UserEmail
	var notActivatedEmailModel *UserEmail
	if ap, ok := emails.NonActivatedEmailData(); ok {
		notActivatedEmailEntity = &ap
		notActivatedEmailModel = EmailToModel(notActivatedEmailEntity)
	}

	return &Emails{
		ActivatedEmail:    activatedEmailModel,
		NotActivatedEmail: notActivatedEmailModel,
	}
}

func EmailToModel(email *emailEntity.UserEmail) *UserEmail {
	var verificationCode *verificationPrimitive.VerificationCode
	var verificationCodeModel *VerificationCode
	if vc, ok := email.VerificationCodeData(); ok {
		verificationCode = &vc
		verificationCodeModel = VerificationToModel(verificationCode)
	}
	return &UserEmail{
		ID:               email.ID().String(),
		Email:            email.Email().String(),
		VerificationCode: verificationCodeModel,
		CreatedAt:        email.CreatedAt().UnixNano(),
	}
}

func getEmailEntity(repoEmail *UserEmail) (*emailEntity.UserEmail, error) {
	emailID, err := emailEntity.EmailIDFrom(repoEmail.ID)
	if err != nil {
		return nil, err
	}

	currentEmail, err := emailPrimitive.EmailFrom(repoEmail.Email)
	if err != nil {
		return nil, err
	}

	var verificationCode *verificationPrimitive.VerificationCode
	if repoEmail.VerificationCode != nil {
		verificationCode, err = VerificationToEntity(repoEmail.VerificationCode)
		if err != nil {
			return nil, err
		}
	}
	emailCreatedAt := commonTime.FromUnixNano(repoEmail.CreatedAt)

	email, err := emailEntity.NewBuilder().
		ID(emailID).
		Email(currentEmail).
		VerificationCode(verificationCode).
		CreatedAt(emailCreatedAt).
		Build()

	return email, nil
}
