package entityStub

import (
	agreementEntity "dev-knowledge/domain/entity/agreement"
	commonTime "dev-knowledge/infrastructure/tools/time"
)

func GetEmptyAgreement() *agreementEntity.Agreement {
	agreement, err := agreementEntity.NewAgreement()
	if err != nil {
		panic(err)
	}

	return agreement
}

func GetAgreement() *agreementEntity.Agreement {
	agreement, err := agreementEntity.NewBuilder().
		Accepted(true).
		AcceptedDate(commonTime.Now()).
		Build()
	if err != nil {
		panic(err)
	}

	return agreement
}
