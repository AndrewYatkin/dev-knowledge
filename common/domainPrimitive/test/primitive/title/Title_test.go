package titlePrimitiveTest

import (
	titlePrimitive "dev-knowledge/common/domainPrimitive/primitive/title"
	commonTesting "dev-knowledge/infrastructure/testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TitleShould struct {
	suite.Suite
}

func TestTitleShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(TitleShould))
}

func (s *TitleShould) TestTitleFrom_ParamIsEmpty_ReturnError() {
	title, err := titlePrimitive.TitleFrom("")

	assert.Empty(s.T(), title)
	assert.ErrorIs(s.T(), err, titlePrimitive.ErrTitleIsEmpty)
}

func (s *TitleShould) TestTitleFrom_TextTooLong_ReturnError() {
	longText := commonTesting.RandomStr(titlePrimitive.TitleMaxLength + 1)

	title, err := titlePrimitive.TitleFrom(longText)

	assert.Empty(s.T(), title)
	assert.ErrorIs(s.T(), err, titlePrimitive.ErrTitleTooLong)
}

func (s *TitleShould) TestTitleFrom_ValidParam_ReturnTitle() {
	expectedTitleStr := commonTesting.RandomDefaultStr()

	title, err := titlePrimitive.TitleFrom(expectedTitleStr)

	require.NoError(s.T(), err)
	assert.Equal(s.T(), expectedTitleStr, title.String())
}
