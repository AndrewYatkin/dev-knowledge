package commonTesting

import (
	"dev-knowledge/infrastructure/errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func AssertErrors(t *testing.T, err error, expectedErrors []error) {
	require.Error(t, err)

	errs, ok := err.(*errors.Errors)
	require.True(t, ok)

	require.Equal(t, len(expectedErrors), errs.Size())
	for _, expectedError := range expectedErrors {
		assert.True(t, errs.Contains(expectedError), fmt.Sprintf("fail assert - missing expected error = %q", expectedError))
	}
}

func AssertErrorCodes(t *testing.T, err error, expectedErrorCodes []errors.ErrorCode) {
	require.Error(t, err)

	errs, ok := err.(*errors.Errors)
	require.True(t, ok)

	actualCodes := map[string]interface{}{}
	for _, actualErr := range errs.ToArray() {
		actualCodes[actualErr.Code().String()] = &struct{}{}
	}

	require.EqualValues(t, len(expectedErrorCodes), errs.Size())
	for _, expectedCode := range expectedErrorCodes {
		_, ok := actualCodes[expectedCode.String()]
		assert.True(t, ok, fmt.Sprintf("fail assert - missing expected error code = %q", expectedCode))
	}
}
