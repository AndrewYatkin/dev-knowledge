package generator

import (
	"dev-knowledge/common/domainPrimitive/primitive/generator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UuidGeneratorShould struct {
	suite.Suite
}

func TestUuidGeneratorShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(UuidGeneratorShould))
}

func (s *UuidGeneratorShould) TestGenerateUUID_ReturnUUID() {
	uuid := generator.GenerateUUID()

	assert.NotEmpty(s.T(), uuid)
}
