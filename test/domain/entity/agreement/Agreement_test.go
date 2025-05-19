package agreement

import (
	agreementEntity "dev-knowledge/domain/entity/agreement"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type AgreementShould struct {
	suite.Suite
}

func TestAgreementShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(AgreementShould))
}

func (s *AgreementShould) TestAccept_IsAccepted() {
	actualAgreement, err := agreementEntity.NewAgreement()

	assert.Nil(s.T(), err)
	assert.False(s.T(), actualAgreement.IsAccepted())
	assert.Nil(s.T(), actualAgreement.AcceptedDate())

	actualAgreement.Accept()

	assert.True(s.T(), actualAgreement.IsAccepted())
	assert.NotNil(s.T(), actualAgreement.AcceptedDate())

	time.Sleep(100)
	acceptedDate1 := actualAgreement.AcceptedDate()
	actualAgreement.Accept()

	assert.True(s.T(), actualAgreement.IsAccepted())
	assert.NotNil(s.T(), actualAgreement.AcceptedDate())
	assert.NotEqual(s.T(), acceptedDate1.UnixNano(), actualAgreement.AcceptedDate().UnixNano())
}
