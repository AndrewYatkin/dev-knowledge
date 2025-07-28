package entityStub

import (
	agreementEntity2 "dev-knowledge/domain/entity/user/agreement"
	commonTime "dev-knowledge/infrastructure/tools/time"
)

func GetEmptyAgreement() *agreementEntity2.Agreement {
	agreement, err := agreementEntity2.NewAgreement()
	if err != nil {
		panic(err)
	}

	return agreement
}

func GetAgreement() *agreementEntity2.Agreement {
	agreement, err := agreementEntity2.NewBuilder().
		Accepted(true).
		AcceptedDate(commonTime.Now()).
		Build()
	if err != nil {
		panic(err)
	}

	return agreement
}
