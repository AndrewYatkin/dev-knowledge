package userRepoModel

import (
	agreementEntity "dev-knowledge/domain/entity/user/agreement"
	commonTime "dev-knowledge/infrastructure/tools/time"
)

type Agreement struct {
	Accepted     bool  `bson:"accepted"`
	AcceptedDate int64 `bson:"accepted_date"`
}

func AgreementToEntity(agreement *Agreement) (*agreementEntity.Agreement, error) {
	var acceptedDate *commonTime.Time
	if agreement.Accepted {
		acceptedDate = commonTime.FromUnixNano(agreement.AcceptedDate)
	}
	agrEnt, err := agreementEntity.NewBuilder().
		Accepted(agreement.Accepted).
		AcceptedDate(acceptedDate).
		Build()
	if err != nil {
		return nil, err
	}

	return agrEnt, nil
}

func AgreementToModel(agreement *agreementEntity.Agreement) *Agreement {
	if agreement.AcceptedDate() == nil {
		return &Agreement{
			Accepted:     agreement.IsAccepted(),
			AcceptedDate: 0,
		}
	}
	return &Agreement{
		Accepted:     agreement.IsAccepted(),
		AcceptedDate: agreement.AcceptedDate().UnixNano(),
	}
}
