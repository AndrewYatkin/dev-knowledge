package agreement

import (
	agreementEntity2 "dev-knowledge/domain/entity/user/agreement"
	"dev-knowledge/infrastructure/errors"
	commonTime "dev-knowledge/infrastructure/tools/time"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AgreementBuilderShould struct {
	suite.Suite
}

func TestAgreementBuilderShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(AgreementBuilderShould))
}

func (s *AgreementBuilderShould) TestBuild_WithoutParams_ReturnWithDefaultValues() {
	actualAgreement, err := agreementEntity2.NewBuilder().Build()

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), actualAgreement)
	assert.False(s.T(), actualAgreement.IsAccepted())
	assert.Nil(s.T(), actualAgreement.AcceptedDate())
}

func (s *AgreementBuilderShould) TestBuild_UnsupportedDateForNotAccepted_ReturnError() {
	expectedAcceptedDate := commonTime.FromUnixNano(999987654321)

	actualAgreement, err := agreementEntity2.NewBuilder().
		Accepted(false).
		AcceptedDate(expectedAcceptedDate).
		Build()

	assert.Nil(s.T(), actualAgreement)
	assert.NotNil(s.T(), err)

	errValue, ok := err.(*errors.Errors)
	assert.True(s.T(), ok)
	assert.EqualValues(s.T(), 1, errValue.Size())
	assert.True(s.T(), errValue.Contains(agreementEntity2.ErrUnsupportedDateForNotAccepted))
}

func (s *AgreementBuilderShould) TestBuild_AcceptWithoutDate_ReturnError() {
	actualAgreement, err := agreementEntity2.NewBuilder().
		Accepted(true).
		AcceptedDate(nil).
		Build()

	assert.Nil(s.T(), actualAgreement)
	assert.NotNil(s.T(), err)

	errValue, ok := err.(*errors.Errors)
	assert.True(s.T(), ok)
	assert.EqualValues(s.T(), 1, errValue.Size())
	assert.True(s.T(), errValue.Contains(agreementEntity2.ErrAcceptedDateIsRequired))
}

func (s *AgreementBuilderShould) TestBuild_AllParams_ReturnAgreement() {
	expectedAcceptedDate := commonTime.FromUnixNano(999987654321)

	actualAgreement, err := agreementEntity2.NewBuilder().
		Accepted(true).
		AcceptedDate(expectedAcceptedDate).
		Build()

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), actualAgreement)
	assert.True(s.T(), actualAgreement.IsAccepted())
	assert.Equal(s.T(), expectedAcceptedDate, actualAgreement.AcceptedDate())
}
