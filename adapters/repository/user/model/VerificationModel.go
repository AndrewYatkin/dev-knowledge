package userRepoModel

import (
	verificationPrimitive "dev-knowledge/common/domainPrimitive/primitive/verification"
	commonTime "dev-knowledge/infrastructure/tools/time"
)

type VerificationCode struct {
	Code      string `bson:"code"`
	CreatedAt int64  `bson:"created_at"`
	ExpireAt  int64  `bson:"expire_at"`
}

func VerificationToEntity(verificationCode *VerificationCode) (*verificationPrimitive.VerificationCode, error) {
	activatedExpireAt := commonTime.FromUnixNano(verificationCode.ExpireAt)
	activatedCreatedAt := commonTime.FromUnixNano(verificationCode.CreatedAt)
	activatedVerificationCode, err := verificationPrimitive.NewBuilder().
		Code(verificationCode.Code).
		ExpireAt(activatedExpireAt).
		CreatedAt(activatedCreatedAt).
		Build()
	if err != nil {
		return nil, err
	}
	return activatedVerificationCode, nil
}
func VerificationToModel(code *verificationPrimitive.VerificationCode) *VerificationCode {
	return &VerificationCode{
		Code:      code.Code(),
		CreatedAt: code.CreatedAt().UnixNano(),
		ExpireAt:  code.ExpireAt().UnixNano(),
	}
}
