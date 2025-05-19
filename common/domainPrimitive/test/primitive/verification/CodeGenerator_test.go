package verification

import (
	verificationPrimitive "dev-knowledge/common/domainPrimitive/primitive/verification"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CodeGeneratorShould struct {
	suite.Suite
}

func TestCodeGeneratorShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(CodeGeneratorShould))
}

func (s *CodeGeneratorShould) TestGenerateCode_ReturnCodeNoErr() {
	expectedCodeLen := 6
	actualCode, err := verificationPrimitive.GenerateCode()

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), actualCode)
	assert.Equal(s.T(), expectedCodeLen, len(actualCode))
}
